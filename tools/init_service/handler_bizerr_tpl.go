package main

const handlerBizerrTpl = `package handler

import "github.com/aiagt/aiagt/common/bizerr"

const (
	ServiceName = "{{ .Service.Snake }}"
{{ range $index, $handler := .Handlers }}
	bizCode{{ $handler.Camel }} = {{ $index }}{{ end }}
)

var ({{ range .Handlers }}
	biz{{ .Camel }} *bizerr.Biz{{ end }}
)

func initServiceBusiness(serviceCode int) {
	baseCode := (serviceCode + 100) * 100
{{ range .Handlers }}
	biz{{ .Camel }} = bizerr.NewBiz(ServiceName, "{{ .Snake }}", baseCode+bizCode{{ .Camel }}){{ end }}
}
`

var HandlerBizerrTpl = NewTemplate("handler.bizerr", handlerBizerrTpl, true)
