package handler

import (
	"errors"
	"fmt"
	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
	"time"

	"github.com/aiagt/aiagt/apps/model/mapper"
	"github.com/aiagt/aiagt/common/ctxutil"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/pkg/closer"
)

func (s *ModelServiceImpl) Chat(req *modelsvc.ChatReq, stream modelsvc.ModelService_ChatServer) (err error) {
	ctx := ctxutil.ApplySpan(stream.Context())

	klog.CtxInfof(ctx, "chat req %v", utils.Pretty(req, 1<<10))

	model, err := s.modelDao.GetByID(ctx, req.ModelId)
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "get model by id failed")
	}

	chatReq := mapper.NewOpenAIGoRequest(req.OpenaiReq, model.ModelKey)

	ok, err := s.callTokenCache.Decr(ctx, req.Token)
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "call token decr failed")
	}

	if !ok {
		return bizChat.NewCodeErr(11, errors.New("call limit reached")).Log(ctx, "call limit reached")
	}

	start := time.Now()
	klog.CtxInfof(ctx, "create chat complation starting")
	chatStream, err := s.openaiCli.CreateChatCompletionStream(ctx, *chatReq)
	klog.CtxInfof(ctx, "create chat complation time consuming: %.2fs", float64(time.Since(start).Milliseconds())/float64(1000))
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "create chat completion stream failed")
	}
	defer closer.Close(chatStream)

	for {
		r, err := chatStream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println()
			return nil
		}

		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "chat stream recv failed")
		}

		fmt.Print(utils.First(r.Choices).Delta.Content)

		openaiResp := mapper.NewOpenAIResponse(&r)

		err = stream.Send(&modelsvc.ChatResp{
			OpenaiResp: openaiResp,
		})
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "chat stream send failed")
		}
	}
}
