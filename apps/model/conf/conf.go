package conf

import (
	"github.com/aiagt/aiagt/common/confutil"
	"path/filepath"

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
}

type OpenAI struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}
