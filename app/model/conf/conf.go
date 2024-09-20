package conf

import (
	"path/filepath"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	ktconf.LoadFiles(conf,
		filepath.Join("conf", "conf.yaml"),
		filepath.Join("app", "model", "conf", "conf.yaml"),
		filepath.Join("conf", "conf-local.yaml"),
		filepath.Join("app", "model", "conf", "conf-local.yaml"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf

	OpenAI OpenAI `yaml:"openai"`
}

type OpenAI struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}
