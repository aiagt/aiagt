package main

const confTpl = `package conf

import (
	"path/filepath"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	confutil.LoadConf(conf,
		filepath.Join("conf"),
		filepath.Join("apps", "{{ .Service.Name }}", "conf"),
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
