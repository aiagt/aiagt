package handler

import (
	"encoding/json"
	"errors"
	"github.com/aiagt/aiagt/pkg/closer"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"

	"github.com/aiagt/aiagt/app/model/mapper"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

func (s *ModelServiceImpl) Chat(req *modelsvc.ChatReq, stream modelsvc.ModelService_ChatServer) (err error) {
	ctx := stream.Context()

	chatReq := mapper.NewOpenAIGoRequest(req.OpenaiReq)
	chatReq.Model = "gpt-3.5-turbo-0125"

	ok, err := s.callTokenCache.Decr(ctx, req.Token)
	if err != nil {
		return bizChat.NewErr(err).Log("call token decr failed")
	}

	if !ok {
		return bizChat.NewCodeErr(11, errors.New("call limit reached")).Log("call limit reached")
	}

	reqJSON, _ := json.MarshalIndent(chatReq, "", "  ")
	klog.Info("chatReq: ", string(reqJSON))

	chatStream, err := s.openaiCli.CreateChatCompletionStream(ctx, *chatReq)
	if err != nil {
		return bizChat.NewErr(err).Log("create chat completion stream failed")
	}
	defer closer.Close(chatStream)

	for {
		r, err := chatStream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return bizChat.NewErr(err).Log("chat stream recv failed")
		}

		s, _ := json.MarshalIndent(r.Choices[0], "", "  ")
		klog.Info(string(s))

		openaiResp := mapper.NewOpenAIResponse(&r)

		err = stream.Send(&modelsvc.ChatResp{
			OpenaiResp: openaiResp,
		})
		if err != nil {
			return bizChat.NewErr(err).Log("chat stream send failed")
		}
	}
}
