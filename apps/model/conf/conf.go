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
		filepath.Join("apps", "model", "conf"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf

	OpenAI  OpenAI  `yaml:"openai"`
	Metrics Metrics `yaml:"metrics"`
	Tracing Tracing `yaml:"tracing"`
}

type OpenAI struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}

type Tracing struct {
	ExportAddr string `yaml:"export_addr"`
}
