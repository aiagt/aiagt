package main

const confYamlTpl = `server:
  name: {{ .Service.Name }}
  address: ":8888"

config_center:
  port: 8848

registry:
  address:
    - ":8848"

db:
  dsn: "root:123456@tcp(127.0.0.1:3306)/aiagt?charset=utf8mb4&parseTime=True&loc=Local"

redis:
  address: "127.0.0.1:6379"
`

var ConfYamlTpl = NewTemplate("conf.yaml", confYamlTpl, false)
