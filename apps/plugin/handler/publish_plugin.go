package handler

import (
	"context"
	"time"

	"github.com/aiagt/aiagt/apps/plugin/model"
	"github.com/aiagt/aiagt/common/bizerr"
	"github.com/aiagt/aiagt/common/ctxutil"
	"github.com/pkg/errors"

	base "github.com/aiagt/aiagt/kitex_gen/base"
)

// PublishPlugin implements the PluginServiceImpl interface.
func (s *PluginServiceImpl) PublishPlugin(ctx context.Context, req *base.IDReq) (resp *base.Empty, err error) {
	plugin, err := s.pluginDao.GetByID(ctx, req.Id)
	if err != nil {
		return nil, bizPublishPlugin.NewErr(err).Log(ctx, "get plugin by id failed")
	}

	if ctxutil.Forbidden(ctx, plugin.AuthorID) {
		return nil, bizPublishPlugin.CodeErr(bizerr.ErrCodeForbidden).Log(ctx, "forbidden")
	}

	tools, err := s.toolDao.GetByPluginID(ctx, req.Id)
	if err != nil {
		return nil, bizPublishPlugin.NewErr(err).Log(ctx, "get tools by plugin id failed")
	}

	for _, tool := range tools {
		// if the test time is before the update time,
		// it means that there are no tests after the update and cannot be published
		if tool.TestedAt == nil || tool.TestedAt.Before(tool.UpdatedAt) {
			return nil, bizPublishPlugin.NewCodeErr(11,
				errors.New("plugin tools not completed testing")).Log(ctx, "plugin tools not completed testing")
		}
	}

	now := time.Now()

	err = s.pluginDao.Update(ctx, req.Id, &model.PluginOptional{PublishedAt: &now})
	if err != nil {
		return nil, bizPublishPlugin.NewErr(err).Log(ctx, "update plugin failed")
	}

	return
}
