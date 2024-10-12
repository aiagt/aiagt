package clientsuite

import (
	"github.com/aiagt/aiagt/common/kitex/clientsuite/middleware"
	ktconf "github.com/aiagt/kitextool/conf"
	ktresolver "github.com/aiagt/kitextool/option/client/resolver"
	ktclient "github.com/aiagt/kitextool/suite/client"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type ClientSuite struct {
	opts []client.Option
}

func (s *ClientSuite) Options() []client.Option {
	return s.opts
}

func NewClientSuite(conf *ktconf.MultiClientConf, svc string) *ClientSuite {
	var opts []client.Option

	opts = append(opts, client.WithTransportProtocol(transport.TTHeaderFramed))
	opts = append(opts, client.WithMetaHandler(transmeta.ClientTTHeaderHandler))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: svc}))

	m := middleware.NewMiddleware()
	opts = append(opts, m.Middlewares()...)

	opts = append(opts, client.WithSuite(ktclient.NewKitexToolSuite(
		conf.GetClientConf(svc),
		ktresolver.WithResolver(ktresolver.NewConsulResolver))))

	opts = append(opts, client.WithSuite(tracing.NewClientSuite()))

	return &ClientSuite{opts: opts}
}
