package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aiagt/aiagt/common/closer"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"

	"github.com/aiagt/aiagt/app/model/mapper"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

func (s *ModelServiceImpl) Chat(req *modelsvc.ChatReq, stream modelsvc.ModelService_ChatServer) (err error) {
	chatReq := mapper.NewOpenAIGoRequest(req.OpenaiReq)

	chatStream, err := s.openaiCli.CreateChatCompletionStream(context.Background(), *chatReq)
	if err != nil {
		return bizChat.NewErr(err)
	}
	defer closer.Close(chatStream)

	for {
		r, err := chatStream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("EOF")
			return nil
		}
		if err != nil {
			return bizChat.NewErr(err)
		}

		s, _ := json.Marshal(r)
		klog.Infof("[RECV] %s", s)

		err = stream.Send(&modelsvc.ChatResp{
			OpenaiResp: mapper.NewOpenAIResponse(&r),
		})
		if err != nil {
			return bizChat.NewErr(err)
		}
	}
}
