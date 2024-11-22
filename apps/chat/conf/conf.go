package conf

import (
	"path/filepath"

	"github.com/aiagt/aiagt/common/confutil"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	confutil.LoadConf(conf,
		filepath.Join("conf"),
		filepath.Join("apps", "chat", "conf"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf

	Metrics Metrics `yaml:"metrics"`
	Tracing Tracing `yaml:"tracing"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}

type Tracing struct {
	ExportAddr string `yaml:"export_addr"`
}
