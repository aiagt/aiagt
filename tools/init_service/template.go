package main

import (
	"strings"
	"text/template"
)

type Template struct {
	tpl     *template.Template
	rewrite bool
}

func NewTemplate(name, tpl string, rewrite bool) *Template {
	tpl = buildSymbol(tpl)
	return &Template{tpl: template.Must(template.New(name).Parse(tpl)), rewrite: rewrite}
}

type symbol [2]string

var BackQuote = symbol{"$0$", "`"}

func buildSymbol(tpl string) string {
	tpl = strings.ReplaceAll(tpl, BackQuote[0], BackQuote[1])
	return tpl
}
