package conf

import (
	"path/filepath"
	"time"

	"github.com/aiagt/aiagt/common/confutil"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	confutil.LoadConf(conf,
		filepath.Join("conf"),
		filepath.Join("apps", "user", "conf"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf

	Email   Email   `yaml:"email"`
	Auth    Auth    `yaml:"auth"`
	Metrics Metrics `yaml:"metrics"`
	Tracing Tracing `yaml:"tracing"`
}

type Email struct {
	SmtpAddr     string `yaml:"smtp_addr"`
	SmtpHost     string `yaml:"smtp_host"`
	EmailFrom    string `yaml:"email_from"`
	EmailAddress string `yaml:"email_address"`
	Auth         string `yaml:"auth"`
}

type Auth struct {
	EncryptSalt   string        `yaml:"encrypt_salt"`
	SnowflakeNode int           `yaml:"snowflake_node"`
	JWTKey        string        `yaml:"jwt_key"`
	JWTExpire     time.Duration `yaml:"jwt_expire"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}

type Tracing struct {
	ExportAddr string `yaml:"export_addr"`
}
