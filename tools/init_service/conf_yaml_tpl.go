package main

import (
	"text/template"
)

const confYamlTpl = `server:
  name: {{ .ServiceName }}
  address: ":8888"

config_center:
  port: 8848

registry:
  address:
    - ":8848"

db:
  dsn: "root:root@tcp(127.0.0.1:3306)/echo?charset=utf8mb4&parseTime=True&loc=Local"

redis:
  address: "127.0.0.1:26379"
`

var ConfYamlTpl = template.Must(template.New("conf.yaml").Parse(confYamlTpl))
