package handler

import (
	"errors"
	"io"

	"github.com/aiagt/aiagt/common/ctxutil"
	"go.opentelemetry.io/otel/trace"

	"github.com/aiagt/aiagt/apps/model/mapper"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/pkg/closer"
)

func (s *ModelServiceImpl) Chat(req *modelsvc.ChatReq, stream modelsvc.ModelService_ChatServer) (err error) {
	ctx := stream.Context()
	ctx = trace.ContextWithSpan(ctx, ctxutil.Span(ctx))

	model, err := s.modelDao.GetByID(ctx, req.ModelId)
	if err != nil {
		return bizChat.CallErr(err).Log(ctx, "get model by id failed")
	}

	chatReq := mapper.NewOpenAIGoRequest(req.OpenaiReq, model.ModelKey)

	ok, err := s.callTokenCache.Decr(ctx, req.Token)
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "call token decr failed")
	}

	if !ok {
		return bizChat.NewCodeErr(11, errors.New("call limit reached")).Log(ctx, "call limit reached")
	}

	chatStream, err := s.openaiCli.CreateChatCompletionStream(ctx, *chatReq)
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "create chat completion stream failed")
	}
	defer closer.Close(chatStream)

	for {
		r, err := chatStream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "chat stream recv failed")
		}

		openaiResp := mapper.NewOpenAIResponse(&r)

		err = stream.Send(&modelsvc.ChatResp{
			OpenaiResp: openaiResp,
		})
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "chat stream send failed")
		}
	}
}
