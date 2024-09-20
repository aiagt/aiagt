package main

import (
	"flag"
	"fmt"
	"github.com/aiagt/aiagt/pkg/closer"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aiagt/aiagt/tools/utils/goparser"
	"github.com/aiagt/aiagt/tools/utils/logger"
	"github.com/aiagt/aiagt/tools/utils/multi_error"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

func main() {
	var (
		serviceName      string
		servicePath      string
		models           flagSlice[string]
		skipDefaultModel bool
	)

	flag.StringVar(&serviceName, "service_name", "", "service name")
	flag.StringVar(&servicePath, "service_path", "", "path to service directory")
	flag.Var(&models, "model", "model name")
	flag.BoolVar(&skipDefaultModel, "skip_default_model", false, "skip default model")
	flag.Parse()

	if serviceName == "" {
		serviceName = filepath.Base(servicePath)
	}

	if !skipDefaultModel {
		models = append(models, serviceName)
	}

	handlers, err := ParseHandlers(servicePath, serviceName)
	if err != nil {
		logger.Errorf("parse handlers error: %s", err)
	}

	render := NewRender(serviceName, handlers)

	path := func(format string, args ...interface{}) string {
		return filepath.Join(servicePath, fmt.Sprintf(format, args...))
	}

	m := multi_error.NewMultiError2[string, *Template]()
	m.Run(render.RenderTemplate, path("conf/conf.go"), ConfTpl).Expect("Render conf error")
	m.Run(render.RenderTemplate, path("conf/conf.yaml"), ConfYamlTpl).Expect("Render conf.yaml error")
	m.Run(render.RenderTemplate, path("model/base.go"), ModelBaseTpl).Expect("Render model.base error")
	m.Run(render.RenderTemplate, path("handler/handler_bizerr.go"), HandlerBizerrTpl).Expect("Render handler.bizerr error")

	for _, model := range models {
		render.Model = NewName(model)
		m.Run(render.RenderTemplate, path("dal/db/%s.go", model), DalDBTpl).Expect("Render dal.db error")
		m.Run(render.RenderTemplate, path("model/%s.go", model), ModelTpl).Expect("Render model error")
	}
}

type Render struct {
	Service  *Name
	Model    *Name
	Handlers []*Name
}

func NewRender(serviceName string, handlers []string) *Render {
	return &Render{
		Service:  NewName(serviceName),
		Handlers: NewNames(handlers),
	}
}

func (t *Render) RenderTemplate(path string, tpl *Template) error {
	if !tpl.rewrite {
		exists, err := IsExists(path)
		if err != nil {
			return errors.Wrap(err, "IsExists error")
		}

		if exists {
			return nil
		}
	}

	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		return errors.Wrap(err, "MkdirAll error")
	}

	file, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "os.Create error")
	}
	defer closer.Close(file)

	err = tpl.tpl.Execute(file, t)
	if err != nil {
		return errors.Wrap(err, "tpl.Execute error")
	}

	if filepath.Ext(path) == ".go" {
		FormatGoFile(path)
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

func ParseHandlers(servicePath, serviceName string) ([]string, error) {
	handlerFile := filepath.Join(servicePath, "handler", "handler.go")

	exists, err := IsExists(handlerFile)
	if err != nil {
		return nil, errors.Wrap(err, "IsExists error")
	}

	if !exists {
		handlerFile = filepath.Join(servicePath, "handler.go")

		exists, err = IsExists(handlerFile)
		if err != nil {
			return nil, errors.Wrap(err, "IsExists error")
		}

		if !exists {
			return nil, errors.Wrap(os.ErrNotExist, "Not found handler file")
		}
	}

	handlerDir := filepath.Dir(handlerFile)

	structMethods, err := goparser.ParseGoFilesInDir(handlerDir)
	if err != nil {
		return nil, errors.Wrap(err, "ParseGoFilesInDir error")
	}

	methods, ok := structMethods[strcase.ToCamel(serviceName)+"ServiceImpl"]
	if !ok {
		return nil, errors.New("Service implementation not found")
	}

	// keep only public methods
	var result []string
	for _, method := range methods.Methods {
		if len(method) > 0 && method[0] >= 'A' && method[0] <= 'Z' {
			result = append(result, method)
		}
	}

	return result, nil
}

type Name struct {
	Name       string
	Camel      string
	LowerCamel string
	Snake      string
}

func NewName(name string) *Name {
	return &Name{
		Name:       name,
		Camel:      strcase.ToCamel(name),
		LowerCamel: strcase.ToLowerCamel(name),
		Snake:      strcase.ToSnake(name),
	}
}

func NewNames(names []string) []*Name {
	result := make([]*Name, len(names))
	for i, name := range names {
		result[i] = NewName(name)
	}

	return result
}

func FormatGoFile(file string) {
	err := exec.Command("go", "fmt", file).Run()
	if err != nil {
		logger.Warnf("format go file %s error: %v", file, err)
	}
}

type flagSlice[T any] []T

func (s *flagSlice[T]) Set(value T) error {
	*s = append(*s, value)
	return nil
}

func (s *flagSlice[T]) String() string {
	return fmt.Sprint(*s)
}
