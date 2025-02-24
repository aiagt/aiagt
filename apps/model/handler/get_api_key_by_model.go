package handler

import (
	"context"
	"github.com/aiagt/aiagt/apps/model/model"
	"github.com/aiagt/aiagt/common/bizerr"

	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// GetAPIKeyByModel implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) GetAPIKeyByModel(ctx context.Context, req *modelsvc.GetAPIKeyByModelReq) (resp *modelsvc.APIKey, err error) {
	var m *model.Models

	if req.ModelId != nil {
		m, err = s.modelDao.GetByID(ctx, req.GetModelId())
	} else if req.Model != nil {
		m, err = s.modelDao.GetByNameOrKey(ctx, req.GetModel())
	} else {
		return nil, bizGetApikeyByModel.CodeErr(bizerr.ErrCodeBadRequest).Log(ctx, "invalid params")
	}

	if err != nil {
		return nil, bizGetApikeyByModel.NewErr(err).Log(ctx, "get api_key by model error", err)
	}

	apiKey, err := s.apiKeyDao.GetBySourceOrDefault(ctx, m.Source)
	if err != nil {
		return nil, bizGetApikeyByModel.NewErr(err)
	}

	resp = &modelsvc.APIKey{
		Id:      apiKey.ID,
		Source:  apiKey.Source,
		BaseUrl: apiKey.BaseURL,
		ApiKey:  apiKey.APIKey,
	}

	return
}
