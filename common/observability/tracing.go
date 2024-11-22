package observability

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

func InitTracing(dest string, exportAddr string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(dest),
		provider.WithExportEndpoint(exportAddr),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)

	server.RegisterShutdownHook(func() {
		logger.Fatal(p.Shutdown(context.Background()))
	})

	return p
}
