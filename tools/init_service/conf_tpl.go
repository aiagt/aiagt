package main

const confTpl = `package conf

import (
	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf
}
`

var ConfTpl = NewTemplate("conf", confTpl, false)
