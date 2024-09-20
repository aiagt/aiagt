package metahandler

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/utils/contextmap"
)

func NewStreamingMetaHandler() remote.MetaHandler {
	return remote.NewCustomMetaHandler(remote.WithOnReadStream(
		func(ctx context.Context) (context.Context, error) {
			return contextmap.WithContextMap(ctx), nil
		},
	))
}
