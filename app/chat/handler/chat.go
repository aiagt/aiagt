package handler

import (
	"context"
	"github.com/aiagt/aiagt/common/safe"
	"github.com/pkg/errors"
	"io"
	"strings"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/app/chat/mapper"
	"github.com/aiagt/aiagt/app/chat/model"
	"github.com/aiagt/aiagt/common/hash/hmap"
	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
)

func (s *ChatServiceImpl) Chat(req *chatsvc.ChatReq, stream chatsvc.ChatService_ChatServer) (err error) {
	ctx := stream.Context()
	userID := ctxutil.UserID(ctx)

	// get app information
	app, err := s.appCli.GetAppByID(ctx, &base.IDReq{Id: req.AppId})
	if err != nil {
		return bizChat.CallErr(err)
	}

	// verify that the user has access rights to the app
	if app.AuthorId != userID {
		return bizChat.CodeErr(bizerr.ErrCodeForbidden)
	}

	// get message records and conversation information
	var (
		msgs         []*model.Message
		conversation *model.Conversation
	)

	if req.ConversationId != nil {
		msgs, err = s.messageDao.GetByConversationID(ctx, *req.ConversationId)
		if err != nil {
			return bizChat.NewErr(err)
		}

		conversation, err = s.conversationDao.GetByID(ctx, *req.ConversationId)
		if err != nil {
			return bizChat.NewErr(err)
		}
	} else {
		const newConversationTitle = "New Conversation"
		conversation = &model.Conversation{
			Title:  newConversationTitle,
			UserID: userID,
			AppID:  req.AppId,
		}

		err = s.conversationDao.Create(ctx, conversation)
		if err != nil {
			return bizChat.CallErr(err)
		}

		req.ConversationId = &conversation.ID
	}

	msgs = append(msgs, mapper.NewModelChatMessage(*req.ConversationId, req.Messages)...)

	// verify that the user has access rights to the conversation
	if conversation.UserID != userID {
		return bizChat.CodeErr(bizerr.ErrCodeForbidden)
	}

	// generate token for model call
	genTokenResp, err := s.modelCli.GenToken(ctx, &modelsvc.GenTokenReq{
		AppId:          req.AppId,
		ConversationId: *req.ConversationId,
		CallLimit:      1,
	})
	if err != nil {
		return bizChat.CallErr(err)
	}

	return s.chat(ctx, *req.ConversationId, msgs, app, genTokenResp.Token, stream)
}

func (s *ChatServiceImpl) chat(ctx context.Context, conversationID int64, msgs []*model.Message, app *appsvc.App, token string, stream chatsvc.ChatService_ChatServer) (err error) {
	var (
		messages       = mapper.NewOpenAIListMessage(msgs)
		modelConfig    = app.ModelConfig
		functions      = mapper.NewOpenAIListFunctionDefinition(app.Tools)
		toolMap        = hmap.NewMapWithKeyFunc(app.Tools, func(t *pluginsvc.PluginTool) string { return t.Name })
		messageContent strings.Builder
	)

	// call model chat api
	chatStream, err := s.modelStreamCli.Chat(ctx, &modelsvc.ChatReq{
		Token:   token,
		ModelId: app.ModelId,
		OpenaiReq: &openai.ChatCompletionRequest{
			Messages:         messages,
			MaxTokens:        modelConfig.MaxTokens,
			Temperature:      modelConfig.Temperature,
			TopP:             modelConfig.TopP,
			N:                &modelConfig.N,
			Stream:           &modelConfig.Stream,
			Stop:             modelConfig.Stop,
			PresencePenalty:  modelConfig.PresencePenalty,
			ResponseFormat:   modelConfig.ResponseFormat,
			Seed:             modelConfig.Seed,
			FrequencyPenalty: modelConfig.FrequencyPenalty,
			LogitBias:        modelConfig.LogitBias,
			Logprobs:         modelConfig.Logprobs,
			TopLogprobs:      modelConfig.TopLogprobs,
			User:             modelConfig.User,
			Functions:        functions,
			StreamOptions:    modelConfig.StreamOptions,
			// FunctionCall:      nil,
			// Tools:             nil,
			// ToolChoice:        nil,
			// ParallelToolCalls: nil,
		},
	})
	if err != nil {
		return bizChat.CallErr(err)
	}

	// traverse to receive and process result
	for {
		resp, err := chatStream.Recv()
		if err == io.EOF {
			msg := &model.Message{
				Base:           model.Base{},
				MessageContent: model.MessageContent{},
				ConversationID: conversationID,
				Role:           model.MessageRoleAssistant,
			}

			// store message
			err = s.messageDao.Create(ctx, msg)
			if err != nil {
				return bizChat.NewErr(err)
			}

			return nil
		}

		if err != nil {
			return bizChat.NewErr(err)
		}

		var (
			functionCallName      strings.Builder
			functionCallArguments strings.Builder
		)
		// parse each choice
		for i, choice := range resp.OpenaiResp.Choices {
			switch {
			case choice.Delta.FunctionCall != nil:
				functionCall := choice.Delta.FunctionCall

				if i == 0 {
					// send function call message
					err := stream.Send(&chatsvc.ChatResp{
						Messages: []*chatsvc.ChatRespMessage{{
							Role: chatsvc.MessageRole_ASSISTANT,
							Content: &chatsvc.MessageContent{
								Type: chatsvc.MessageType_FUNCTION_CALL,
								Content: &chatsvc.MessageContentValue{FuncCall: &chatsvc.MessageContentValueFuncCall{
									Name:      functionCall.GetName(),
									Arguments: functionCall.GetArguments(),
								}},
							}}},
						ConversationId: conversationID,
					})
					if err != nil {
						return bizChat.CallErr(err)
					}
				}

				functionCallName.WriteString(functionCall.GetName())
				functionCallArguments.WriteString(functionCall.GetArguments())
			case choice.Delta.Content != nil:
				content := choice.Delta.Content
				messageContent.WriteString(*content)

				// send text message
				err = stream.Send(&chatsvc.ChatResp{
					Messages: []*chatsvc.ChatRespMessage{{
						Role: chatsvc.MessageRole_ASSISTANT,
						Content: &chatsvc.MessageContent{
							Type:    chatsvc.MessageType_TEXT,
							Content: &chatsvc.MessageContentValue{Text: &chatsvc.MessageContentValueText{Text: *content}},
						}}},
					ConversationId: conversationID,
				})
				if err != nil {
					return bizChat.CallErr(err)
				}
			case choice.FinishReason == "function_call":
				functionCall := &openai.FunctionCall{
					Name:      safe.Pointer(functionCallName.String()),
					Arguments: safe.Pointer(functionCallArguments.String()),
				}
				tool, ok := toolMap[functionCall.GetName()]
				if !ok {
					return bizChat.NewErr(errors.New("plugin tool not found"))
				}
				return s.handleFunctionCall(ctx, functionCall, conversationID, tool, stream)
			}
		}
	}
}

func (s *ChatServiceImpl) handleFunctionCall(ctx context.Context, functionCall *openai.FunctionCall, conversationID int64, tool *pluginsvc.PluginTool, stream chatsvc.ChatService_ChatServer) error {
	// get user secrets
	const maxSecrets = 100
	listSecretResp, err := s.userCli.ListSecret(ctx, &usersvc.ListSecretReq{
		Pagination: &base.PaginationReq{PageSize: maxSecrets},
		PluginId:   &tool.PluginId,
	})
	if err != nil {
		return bizChat.CallErr(err)
	}

	secretMap := hmap.NewMapWithFuncs(listSecretResp.Secrets,
		func(t *usersvc.Secret) string { return t.Name },
		func(t *usersvc.Secret) string { return t.Value })

	// call plugin tool
	callResp, err := s.pluginCli.CallPluginTool(ctx, &pluginsvc.CallPluginToolReq{
		PluginId: tool.PluginId,
		ToolId:   tool.Id,
		Secrets:  secretMap,
		Request:  []byte(*functionCall.Arguments),
	})
	if err != nil {
		return bizChat.CallErr(err)
	}

	// store the result of the call
	msg := &model.Message{
		MessageContent: model.MessageContent{
			Type: model.MessageTypeFunction,
			Content: model.MessageContentValue{Func: &model.MessageContentValueFunc{
				Name:    *functionCall.Name,
				Content: string(callResp.Response),
			}}},
		ConversationID: conversationID,
		Role:           model.MessageRoleFunction,
	}
	err = s.messageDao.Create(ctx, msg)
	if err != nil {
		return bizChat.NewErr(err)
	}

	// send function result message
	err = stream.Send(&chatsvc.ChatResp{
		Messages: []*chatsvc.ChatRespMessage{{
			Role: chatsvc.MessageRole_FUNCTION,
			Content: &chatsvc.MessageContent{
				Type: chatsvc.MessageType_FUNCTION,
				Content: &chatsvc.MessageContentValue{Func: &chatsvc.MessageContentValueFunc{
					Name:    *functionCall.Name,
					Content: string(callResp.Response),
				}},
			}}},
		ConversationId: conversationID,
	})
	if err != nil {
		return bizChat.CallErr(err)
	}

	return nil
}
