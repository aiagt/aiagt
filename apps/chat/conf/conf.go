package conf

import (
	"path/filepath"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	ktconf.LoadFiles(conf,
		filepath.Join("conf", "conf.yaml"),
		filepath.Join("apps", "chat", "conf", "conf.yaml"),
		"/Users/user/Code/aiagt/apps/chat/conf/conf.yaml",
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf

	Metrics Metrics `yaml:"metrics"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}
