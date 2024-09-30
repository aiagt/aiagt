package handler

import (
	"context"
	"time"

	"github.com/aiagt/aiagt/apps/model/dal/cache"
	"github.com/aiagt/aiagt/apps/user/pkg/encrypt"
	"github.com/aiagt/aiagt/common/baseutil"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
	"github.com/google/uuid"
)

// GenToken implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) GenToken(ctx context.Context, req *modelsvc.GenTokenReq) (resp *modelsvc.GenTokenResp, err error) {
	val := &cache.CallTokenValue{
		AppID:          req.AppId,
		PluginID:       req.PluginId,
		ConversationID: req.ConversationId,
		CallLimit:      req.CallLimit,
	}

	token := encrypt.Encrypt(uuid.New().String())

	err = s.callTokenCache.Set(ctx, token, val)
	if err != nil {
		return nil, bizGenToken.NewErr(err).Log("set call token cache failed")
	}

	resp = &modelsvc.GenTokenResp{
		Token:     token,
		ExpiredAt: baseutil.NewBaseTime(time.Now().Add(10 * time.Minute)),
	}

	return
}
