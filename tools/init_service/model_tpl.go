package main

const modelTpl = `package model

type {{ .CamelServiceName }} struct {
	Base
}
`

var ModelTpl = NewTemplate("model", modelTpl, false)
