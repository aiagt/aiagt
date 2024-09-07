package main

const confTpl = `package conf

import (
	"path/filepath"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	ktconf.LoadFiles(conf,
		filepath.Join("conf", "conf.yaml"),
		filepath.Join("app", "{{ .Service.Name }}", "conf", "conf.yaml"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf
}
`

var ConfTpl = NewTemplate("conf", confTpl, false)
