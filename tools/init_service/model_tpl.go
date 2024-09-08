package main

const modelTpl = `package model

type {{ .Model.Camel }} struct {
	Base
}

type {{ .Model.Camel }}Optional struct {
}
`

var ModelTpl = NewTemplate("model", modelTpl, false)
