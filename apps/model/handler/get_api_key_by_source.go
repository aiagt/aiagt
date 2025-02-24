package handler

import (
	"context"
	modelsvc "github.com/aiagt/aiagt/kitex_gen/modelsvc"
)

// GetAPIKeyBySource implements the ModelServiceImpl interface.
func (s *ModelServiceImpl) GetAPIKeyBySource(ctx context.Context, req *modelsvc.GetAPIKeyBySourceReq) (resp *modelsvc.APIKey, err error) {
	apiKey, err := s.apiKeyDao.GetBySourceOrDefault(ctx, req.Source)
	if err != nil {
		return nil, bizGetApikeyBySource.NewErr(err).Log(ctx, "get api_key by source err", err)
	}

	resp = &modelsvc.APIKey{
		Id:      apiKey.ID,
		Source:  apiKey.Source,
		BaseUrl: apiKey.BaseURL,
		ApiKey:  apiKey.APIKey,
	}

	return
}
