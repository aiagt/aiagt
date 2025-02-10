package conf

import (
	"path/filepath"
	"strings"

	"github.com/aiagt/aiagt/common/confutil"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	confutil.LoadConf(conf,
		filepath.Join("conf"),
		filepath.Join("apps", "model", "conf"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf

	APIKeys `yaml:"api_keys"`
	Metrics Metrics `yaml:"metrics"`
	Tracing Tracing `yaml:"tracing"`
}

type APIKeys map[string]APIKey

func (k APIKeys) Default() *APIKey {
	const defaultKey = "default"

	result, ok := k.Get(defaultKey)
	if !ok {
		return &APIKey{}
	}

	return result
}

func (k APIKeys) Get(key string) (*APIKey, bool) {
	result, ok := k[strings.ToLower(key)]
	if !ok {
		return nil, false
	}

	return &result, true
}

func (k APIKeys) GetOrDefault(key string) *APIKey {
	result, ok := k.Get(key)
	if !ok {
		return k.Default()
	}

	return result
}

type APIKey struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}

type Tracing struct {
	ExportAddr string `yaml:"export_addr"`
}
