package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/pkg/errors"

	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"

	"github.com/aiagt/aiagt/apps/chat/mapper"
	"github.com/aiagt/aiagt/apps/chat/model"
	"github.com/aiagt/aiagt/kitex_gen/appsvc"
	"github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"github.com/aiagt/aiagt/kitex_gen/pluginsvc"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
)

func (s *ChatServiceImpl) Chat(req *chatsvc.ChatReq, stream chatsvc.ChatService_ChatServer) (err error) {
	ctx := ctxutil.ApplySpan(stream.Context())
	userID := ctxutil.UserID(ctx)

	user, err := s.userCli.GetUserByID(ctx, &base.IDReq{Id: userID})
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "get user failed")
	}

	// get app information
	getAppResp, err := s.appCli.GetAppByID(ctx, &appsvc.GetAppByIDReq{Id: req.AppId, Unfold: utils.Pointer(true)})
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "get app by id error")
	}
	var app = getAppResp.App

	// verify that the user has access rights to the app
	if app.IsPrivate && app.AuthorId != userID {
		return bizChat.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "user does not have access rights to the app")
	}

	// get message records and conversation information
	var (
		msgs         []*model.Message
		conversation *model.Conversation
	)

	if req.ConversationId != nil {
		msgs, err = s.messageDao.GetByConversationID(ctx, *req.ConversationId)
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "get message by conversation id error")
		}

		conversation, err = s.conversationDao.GetByID(ctx, *req.ConversationId)
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "get conversation by id error")
		}
	} else {
		const defaultConversationTitle = "New Conversation"
		conversation = &model.Conversation{
			Title:  defaultConversationTitle,
			UserID: userID,
			AppID:  req.AppId,
		}

		err = s.conversationDao.Create(ctx, conversation)
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "create conversation error")
		}

		req.ConversationId = &conversation.ID

		wg := new(sync.WaitGroup)
		defer wg.Wait()

		wg.Add(1)
		go func() {
			defer wg.Done()

			msg, _ := json.Marshal(req.Messages)

			const (
				modelGPT35Turbo0125ID = 1
				modelGPT35TurboID     = 2
				modelGPT35Turbo16kID  = 3
				modelGPT4oID          = 4
			)

			var retryModelIDs = []int64{modelGPT35Turbo0125ID, modelGPT4oID}

			// generate title, retry with different model
			for _, modelID := range retryModelIDs {
				ok := s.generateNewTitle(ctx, stream, string(msg), *req.ConversationId, modelID)
				if ok {
					break
				}
			}
		}()
	}

	newMsgs := mapper.NewModelChatMessage(*req.ConversationId, req.Messages)

	if len(newMsgs) > 0 {
		err = s.messageDao.CreateBatch(ctx, newMsgs)
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "create message batch error")
		}

		msgs = append(msgs, newMsgs...)

		err = stream.Send(&chatsvc.ChatResp{
			Messages:       mapper.NewGenListMessage(newMsgs),
			ConversationId: conversation.ID,
		})
		if err != nil {
			return bizChat.CallErr(err).Log(ctx, "send user text message error")
		}
	}

	// verify that the user has access rights to the conversation
	if conversation.UserID != userID {
		return bizChat.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "user does not have access rights to the conversation")
	}

	// generate token for model call
	callToken, err := s.genToken(ctx, req.AppId, *req.ConversationId, 10)
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "generate token error")
	}

	return s.chat(ctx, *req.ConversationId, user, msgs, app, callToken, stream)
}

func (s *ChatServiceImpl) genToken(ctx context.Context, appID, conversationID int64, callLimit int32) (string, error) {
	genTokenResp, err := s.modelCli.GenToken(ctx, &modelsvc.GenTokenReq{
		AppId:          appID,
		ConversationId: conversationID,
		CallLimit:      callLimit,
	})
	if err != nil {
		return "", bizChat.CallErr(err).Log(ctx, "generate token error")
	}

	return genTokenResp.Token, nil
}

func (s *ChatServiceImpl) chat(ctx context.Context, conversationID int64, user *usersvc.User, msgs []*model.Message, app *appsvc.App, token string, stream chatsvc.ChatService_ChatServer) (err error) {
	var (
		messages    = mapper.NewOpenAIListMessage(msgs)
		modelConfig = app.ModelConfig
		functions   = mapper.NewOpenAIListFunctionDefinition(app.Tools)
		toolMap     = hmap.FromSliceEntries(app.Tools, func(t *pluginsvc.PluginTool) (string, *pluginsvc.PluginTool, bool) { return t.Name, t, true })
	)

	// call model chat api
	chatStream, err := s.modelStreamCli.Chat(ctx, &modelsvc.ChatReq{
		Token:   token,
		ModelId: app.ModelId,
		OpenaiReq: &openai.ChatCompletionRequest{
			Messages: append([]*openai.ChatCompletionMessage{{
				Role:    "system",
				Content: utils.Pointer(fmt.Sprintf(`You are an agent on the ai agent platform "Aiagt", your identity information is:\nName: %s,\nDescription: "%s",\nAuthor: %s`, app.Name, app.Description, app.Author.Username)),
			}}, messages...),
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
			User:             &user.Username,
			Functions:        functions,
			StreamOptions:    modelConfig.StreamOptions,
		},
	})
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "chat stream error")
	}

	var (
		functionCallName      strings.Builder
		functionCallArguments strings.Builder
		messageContent        strings.Builder
	)
	// traverse to receive and process result
	for {
		resp, err := chatStream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "receive chat stream error")
		}

		// parse each choice
		for _, choice := range resp.OpenaiResp.Choices {
			switch {
			case choice.Delta.FunctionCall != nil:
				functionCall := choice.Delta.FunctionCall
				functionCallName.WriteString(functionCall.GetName())
				functionCallArguments.WriteString(functionCall.GetArguments())
			case choice.Delta.Content != nil:
				content := choice.Delta.Content
				messageContent.WriteString(*content)

				// send text message
				err = stream.Send(&chatsvc.ChatResp{
					Messages: []*chatsvc.Message{{
						Role: chatsvc.MessageRole_ASSISTANT,
						Content: &chatsvc.MessageContent{
							Type:    chatsvc.MessageType_TEXT,
							Content: &chatsvc.MessageContentValue{Text: &chatsvc.MessageContentValueText{Text: *content}},
						},
					}},
					ConversationId: conversationID,
				})
				if err != nil {
					return bizChat.CallErr(err).Log(ctx, "send text message error")
				}
			case choice.FinishReason == "function_call":
				functionCall := &openai.FunctionCall{
					Name:      utils.Pointer(functionCallName.String()),
					Arguments: utils.Pointer(functionCallArguments.String()),
				}

				// send function call message
				err := stream.Send(&chatsvc.ChatResp{
					Messages: []*chatsvc.Message{{
						Role: chatsvc.MessageRole_ASSISTANT,
						Content: &chatsvc.MessageContent{
							Type: chatsvc.MessageType_FUNCTION_CALL,
							Content: &chatsvc.MessageContentValue{FuncCall: &chatsvc.MessageContentValueFuncCall{
								Name:      functionCall.GetName(),
								Arguments: functionCall.GetArguments(),
							}},
						},
					}},
					ConversationId: conversationID,
				})
				if err != nil {
					return bizChat.CallErr(err).Log(ctx, "send function call message error")
				}

				tool, ok := toolMap[functionCall.GetName()]
				if !ok {
					return bizChat.NewErr(errors.New("plugin tool not found")).Log(ctx, "plugin tool not found")
				}

				// store function call message
				msg := &model.Message{
					MessageContent: model.MessageContent{
						Type: model.MessageTypeFunctionCall,
						Content: model.MessageContentValue{FuncCall: &model.MessageContentValueFuncCall{
							Name:      functionCall.GetName(),
							Arguments: functionCall.GetArguments(),
						}},
					},
					ConversationID: conversationID,
					Role:           model.MessageRoleAssistant,
				}

				err = s.messageDao.Create(ctx, msg)
				if err != nil {
					return bizChat.NewErr(err).Log(ctx, "create function call message error")
				}

				msgs = append(msgs, msg)

				return s.handleFunctionCall(ctx, functionCall, tool, conversationID, user, msgs, app, token, stream)
			case choice.FinishReason == "stop":
				msg := &model.Message{
					MessageContent: model.MessageContent{
						Type: model.MessageTypeText,
						Content: model.MessageContentValue{Text: &model.MessageContentValueText{
							Text: messageContent.String(),
						}},
					},
					ConversationID: conversationID,
					Role:           model.MessageRoleAssistant,
				}

				err = s.messageDao.Create(ctx, msg)
				if err != nil {
					return bizChat.NewErr(err).Log(ctx, "create text message error")
				}

				stopMsg := mapper.NewGenMessage(msg)
				stopMsg.Content.Content.Text = &chatsvc.MessageContentValueText{}

				err := stream.Send(&chatsvc.ChatResp{
					Messages:       []*chatsvc.Message{stopMsg},
					ConversationId: conversationID,
				})
				if err != nil {
					return bizChat.CallErr(err).Log(ctx, "send stop text message error")
				}

				return nil
			}
		}
	}
}

func (s *ChatServiceImpl) handleFunctionCall(ctx context.Context, functionCall *openai.FunctionCall, tool *pluginsvc.PluginTool, conversationID int64, user *usersvc.User, msgs []*model.Message, app *appsvc.App, token string, stream chatsvc.ChatService_ChatServer) error {
	// get user secrets
	const maxSecrets = 100

	listSecretResp, err := s.userCli.ListSecret(ctx, &usersvc.ListSecretReq{
		Pagination: &base.PaginationReq{PageSize: maxSecrets},
		PluginId:   &tool.PluginId,
	})
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "list secret error")
	}

	secretMap := hmap.FromSliceEntries(listSecretResp.Secrets, func(t *usersvc.Secret) (string, string, bool) { return t.Name, t.Value, true })

	// call plugin tool
	callResp, err := s.pluginCli.CallPluginTool(ctx, &pluginsvc.CallPluginToolReq{
		PluginId: tool.PluginId,
		ToolId:   tool.Id,
		Secrets:  secretMap,
		Request:  []byte(*functionCall.Arguments),
	})
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "call plugin tool error")
	}

	// store the result of the call
	msg := &model.Message{
		MessageContent: model.MessageContent{
			Type: model.MessageTypeFunction,
			Content: model.MessageContentValue{Func: &model.MessageContentValueFunc{
				Name:    *functionCall.Name,
				Content: string(callResp.Response),
			}},
		},
		ConversationID: conversationID,
		Role:           model.MessageRoleFunction,
	}

	err = s.messageDao.Create(ctx, msg)
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "create function message error")
	}

	// send function result message
	err = stream.Send(&chatsvc.ChatResp{
		Messages: []*chatsvc.Message{{
			Role: chatsvc.MessageRole_FUNCTION,
			Content: &chatsvc.MessageContent{
				Type: chatsvc.MessageType_FUNCTION,
				Content: &chatsvc.MessageContentValue{Func: &chatsvc.MessageContentValueFunc{
					Name:    *functionCall.Name,
					Content: string(callResp.Response),
				}},
			},
		}},
		ConversationId: conversationID,
	})
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "send function result message error")
	}

	msgs = append(msgs, msg)

	return s.chat(ctx, conversationID, user, msgs, app, token, stream)
}

// the return value determines whether a retry is required
func (s *ChatServiceImpl) generateNewTitle(ctx context.Context, stream chatsvc.ChatService_ChatServer, msg string, conversationID int64, modelId int64) bool {
	if len(msg) == 0 {
		return true
	}

	type responseFormat struct {
		Title string `json:"title"`
	}

	callToken, err := s.genToken(ctx, 0, 0, 1)
	if err != nil {
		return true
	}

	chatStream, err := s.modelStreamCli.Chat(ctx, &modelsvc.ChatReq{
		Token:   callToken,
		ModelId: modelId,
		OpenaiReq: &openai.ChatCompletionRequest{
			Messages: []*openai.ChatCompletionMessage{
				{
					Role:    mapper.NewOpenAIMessageRole(model.MessageRoleSystem),
					Content: utils.Pointer("Create a short title based on the user's conversation content. The number of characters should not exceed 32. The English word should not exceed 5 words. Try to keep it within 4. The language of the title content should be consistent with the language of the conversation content."),
				},
				{
					Role:    mapper.NewOpenAIMessageRole(model.MessageRoleUser),
					Content: &msg,
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatType_JSON_SCHEMA,
				JsonSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:        "TitleResponse",
					Description: utils.Pointer("JSON schema for title response"),
					Schema:      `{ "type": "object", "properties": { "title": { "type": "string", "description": "Conversation title" } }, "required": [ "title" ], "additionalProperties": false }`,
					Strict:      utils.Pointer(true),
				}},
		},
	})

	if err != nil {
		klog.CtxErrorf(ctx, "generate new title request chat err: %v", err)
		return false
	}

	var title bytes.Buffer

	for {
		resp, err := chatStream.Recv()

		if err == io.EOF {
			var result responseFormat

			err = json.Unmarshal(title.Bytes(), &result)
			if err != nil {
				klog.CtxErrorf(ctx, "generate new title parse result err: %v", err)
				return false
			}

			err = s.conversationDao.Update(ctx, conversationID, &model.ConversationOptional{Title: &result.Title})
			if err != nil {
				klog.CtxErrorf(ctx, "save new title err: %v", err)
				return true
			}

			if len(result.Title) > 0 {
				err = stream.Send(&chatsvc.ChatResp{ConversationTitle: &result.Title})

				if err != nil {
					klog.CtxErrorf(ctx, "generate new title send result title err: %v", err)
					return true
				}
			}

			return true
		}

		if err != nil {
			klog.CtxErrorf(ctx, "build new title stream recv err: %v", err)
			return false
		}

		if resp != nil && resp.OpenaiResp != nil && len(resp.OpenaiResp.Choices) > 0 && resp.OpenaiResp.Choices[0].Delta != nil {
			title.WriteString(utils.Value(resp.OpenaiResp.Choices[0].Delta.Content))
		}
	}
}
