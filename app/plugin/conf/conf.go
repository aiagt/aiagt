package conf

import (
	ktconf "github.com/aiagt/kitextool/conf"
	"path/filepath"
)

var conf = new(ServerConf)

func init() {
	ktconf.LoadFiles(conf,
		filepath.Join("conf", "conf.yaml"),
		filepath.Join("app", "plugin", "conf", "conf.yaml"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf
}
