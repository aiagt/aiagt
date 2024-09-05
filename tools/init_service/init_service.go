package main

import (
	"flag"
	"os"
	"path/filepath"
	"text/template"

	"github.com/aiagt/aiagt/tools/utils/closer"
	"github.com/aiagt/aiagt/tools/utils/multi_error"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

func main() {
	var (
		serviceName string
		servicePath string
	)

	flag.StringVar(&serviceName, "service_name", "", "service name")
	flag.StringVar(&servicePath, "service_path", "", "path to service directory")
	flag.Parse()

	if serviceName == "" {
		serviceName = filepath.Base(servicePath)
	}

	render := NewRender(serviceName)

	m := multi_error.NewMultiError2[string, *template.Template]()
	m.Run(render.RenderTemplate, filepath.Join(servicePath, "conf/conf.go"), ConfTpl).
		Expect("Render conf.go error")
	m.Run(render.RenderTemplate, filepath.Join(servicePath, "conf/conf.yaml"), ConfYamlTpl).
		Expect("Render conf.yaml error")
	m.Run(render.RenderTemplate, filepath.Join(servicePath, "dal/db/db.go"), DaoDBTpl).
		Expect("Render db.go error")
}

type Render struct {
	ServiceName      string
	CamelServiceName string
}

func NewRender(serviceName string) *Render {
	return &Render{ServiceName: serviceName, CamelServiceName: strcase.ToCamel(serviceName)}
}

func (t *Render) RenderTemplate(path string, tpl *template.Template) error {
	exists, err := IsExists(path)
	if err != nil {
		return errors.Wrap(err, "IsExists error")
	}

	if exists {
		return nil
	}

	err = os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return errors.Wrap(err, "MkdirAll error")
	}

	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "os.Create error")
	}
	defer closer.Close(file)

	err = tpl.Execute(file, t)
	if err != nil {
		return errors.Wrap(err, "tpl.Execute error")
	}

	return nil
}

func IsExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, errors.Wrap(err, "os.Stat error")
	}

	return true, nil
}
