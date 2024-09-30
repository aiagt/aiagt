package observability

import (
	"github.com/aiagt/aiagt/pkg/logerr"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

var Registry *prometheus.Registry

func InitMetrics(dest, metricsAddr, registryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	r, err := consul.NewConsulRegister(registryAddr)
	logerr.Fatal(err)

	addr, err := net.ResolveTCPAddr("tcp", metricsAddr)
	logerr.Fatal(err)

	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": dest},
	}

	err = r.Register(registryInfo)
	logerr.Fatal(err)

	server.RegisterShutdownHook(func() {
		err = r.Deregister(registryInfo)
		logerr.Fatal(err)
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	go func() {
		_ = http.ListenAndServe(metricsAddr, nil)
	}()

	return r, registryInfo
}
