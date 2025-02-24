package handler

import (
	"errors"
	"fmt"
	"github.com/aiagt/aiagt/apps/model/model"
	"github.com/sashabaranov/go-openai"
	"io"
	"time"

	"github.com/aiagt/aiagt/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/aiagt/aiagt/apps/model/mapper"
	"github.com/aiagt/aiagt/common/ctxutil"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/aiagt/aiagt/pkg/closer"
)

func (s *ModelServiceImpl) Chat(req *modelsvc.ChatReq, stream modelsvc.ModelService_ChatServer) (err error) {
	ctx := ctxutil.ApplySpan(stream.Context())

	klog.CtxInfof(ctx, "chat req %v", utils.Pretty(req, 1<<10))

	var (
		modelKey    = req.OpenaiReq.Model
		modelSource = model.DefaultSource
	)

	if utils.IsZero(modelKey) {
		m, err := s.modelDao.GetByID(ctx, req.ModelId)
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "get model by id failed")
		}

		modelKey = m.ModelKey
		modelSource = m.Source
	}

	chatReq := mapper.NewOpenAIGoRequest(req.OpenaiReq, modelKey)

	ok, err := s.callTokenCache.Decr(ctx, req.Token)
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "call token decr failed")
	}

	if !ok {
		return bizChat.NewCodeErr(11, errors.New("call limit reached")).Log(ctx, "call limit reached")
	}

	// get llm apikey
	apiKey, err := s.apiKeyDao.GetBySourceOrDefault(ctx, modelSource)
	if err != nil {
		return bizChat.NewErr(err).Log(ctx, "get api key failed")
	}

	start := time.Now()

	klog.CtxDebugf(ctx, "create chat complation req: %s", utils.Pretty(chatReq, 1<<10))
	chatStream, err := newOpenaiCli(apiKey).CreateChatCompletionStream(ctx, *chatReq)
	klog.CtxDebugf(ctx, "create chat complation time consuming: %.2fs", float64(time.Since(start).Milliseconds())/float64(1000))

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

		klog.CtxDebugf(ctx, utils.Pretty(r, 1<<10))

		openaiResp := mapper.NewOpenAIResponse(&r)

		err = stream.Send(&modelsvc.ChatResp{
			OpenaiResp: openaiResp,
		})
		if err != nil {
			return bizChat.NewErr(err).Log(ctx, "chat stream send failed")
		}
	}
}

func newOpenaiCli(apiKey *model.ApiKey) *openai.Client {
	config := openai.DefaultConfig(apiKey.APIKey)
	config.BaseURL = apiKey.BaseURL

	return openai.NewClientWithConfig(config)
}
