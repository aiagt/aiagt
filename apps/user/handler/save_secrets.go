package handler

import (
	"context"
	"fmt"
	"github.com/aiagt/aiagt/apps/user/model"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/aiagt/aiagt/pkg/hash/hmap"
	"github.com/aiagt/aiagt/pkg/hash/hset"
	"time"

	base "github.com/aiagt/aiagt/kitex_gen/base"
	usersvc "github.com/aiagt/aiagt/kitex_gen/usersvc"
)

// SaveSecrets implements the UserServiceImpl interface.
func (s *UserServiceImpl) SaveSecrets(ctx context.Context, req *usersvc.SaveSecretReq) (resp *base.Empty, err error) {
	var (
		userID = ctxutil.UserID(ctx)

		pluginIDs = hset.NewSet[int64](0)
		names     = hset.NewSet[string](0)
	)

	for _, item := range req.Secrets {
		pluginIDs.Add(item.PluginId)
		names.Add(item.Name)
	}

	secrets, err := s.secretDao.ListByPluginsAndNames(ctx, pluginIDs.List(), names.List())
	if err != nil {
		return nil, bizSaveSecrets.NewErr(err).Log(ctx, "get secrets by plugins and names error")
	}

	secretsMap := hmap.FromSliceEntries(secrets, func(t *model.Secret) (string, *model.Secret, bool) {
		return fmt.Sprintf("%d-%s", t.PluginID, t.Name), t, true
	})

	for _, item := range req.Secrets {
		var secret *model.Secret

		if t, ok := secretsMap[fmt.Sprintf("%d-%s", item.PluginId, item.Name)]; ok {
			secret = t
		} else {
			secret = &model.Secret{
				Base: model.Base{CreatedAt: time.Now()},
			}
		}

		secret.UserID = userID
		secret.PluginID = item.PluginId
		secret.Name = item.Name
		secret.Value = item.Value
		secret.UpdatedAt = time.Now()

		err = s.secretDao.Save(ctx, secret)
		if err != nil {
			return nil, bizSaveSecrets.NewErr(err).Log(ctx, "save secret error")
		}
	}

	return
}
