package handler

import (
	"github.com/aiagt/aiagt/app/chat/mapper"
	"github.com/aiagt/aiagt/app/chat/model"
	"github.com/aiagt/aiagt/kitex_gen/base"
	chatsvc "github.com/aiagt/aiagt/kitex_gen/chatsvc"
	"github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/kitex_gen/openai"
	"io"
)

func (s *ChatServiceImpl) Chat(req *chatsvc.ChatReq, stream chatsvc.ChatService_ChatServer) (err error) {
	ctx := stream.Context()

	app, err := s.appCli.GetAppByID(ctx, &base.IDReq{Id: req.AppId})
	if err != nil {
		return bizChat.CallErr(err)
	}

	var msgList []*model.Message

	if req.ConversationId != nil {
		msgList, err = s.messageDao.GetByConversationID(ctx, *req.ConversationId)
		if err != nil {
			return bizChat.NewErr(err)
		}
	} else {
		const unknownTitle = "未命名会话"
		conversation := &model.Conversation{Title: unknownTitle}

		err = s.conversationDao.Create(ctx, conversation)
		if err != nil {
			return bizChat.CallErr(err)
		}

		req.ConversationId = &conversation.ID
	}

	genTokenResp, err := s.modelCli.GenToken(ctx, &modelsvc.GenTokenReq{
		AppId:          req.AppId,
		ConversationId: *req.ConversationId,
		CallLimit:      1,
	})
	if err != nil {
		return bizChat.CallErr(err)
	}

	msgList = append(msgList, mapper.NewModelChatMessage(*req.ConversationId, req.Messages)...)

	var (
		messages    = mapper.NewOpenAIListMessage(msgList)
		modelConfig = app.ModelConfig
		functions   = mapper.NewOpenAIListFunctionDefinition(app.Tools)
	)

	chatStream, err := s.modelStreamCli.Chat(ctx, &modelsvc.ChatReq{
		Token:   genTokenResp.Token,
		ModelId: app.ModelId,
		OpenaiReq: &openai.ChatCompletionRequest{
			Messages:         messages,
			Temperature:      modelConfig.Temperature,
			TopP:             modelConfig.TopP,
			N:                &modelConfig.N,
			Stream:           &modelConfig.Stream,
			PresencePenalty:  modelConfig.PresencePenalty,
			ResponseFormat:   modelConfig.ResponseFormat,
			Seed:             modelConfig.Seed,
			FrequencyPenalty: modelConfig.FrequencyPenalty,
			LogitBias:        modelConfig.LogitBias,
			Logprobs:         modelConfig.Logprobs,
			TopLogprobs:      modelConfig.TopLogprobs,
			Functions:        functions,
			StreamOptions:    modelConfig.StreamOptions,
			//MaxTokens:         modelConfig.MaxTokens,
			//Stop:              modelConfig.Stop,
			//User:              modelConfig.User,
			//FunctionCall:      nil,
			//Tools:             nil,
			//ToolChoice:        nil,
			//ParallelToolCalls: nil,
		},
	})
	if err != nil {
		return bizChat.NewErr(err)
	}

	for {
		resp, err := chatStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return bizChat.NewErr(err)
		}

		for _, choice := range resp.OpenaiResp.Choices {
			if choice.Message.FunctionCall != nil {
				//choice.Message.FunctionCall.
			}
		}
	}

	return
}
