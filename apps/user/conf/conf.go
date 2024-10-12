package conf

import (
	"path/filepath"
	"time"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	ktconf.LoadFiles(conf,
		filepath.Join("conf", "conf.yaml"),
		filepath.Join("apps", "user", "conf", "conf.yaml"),
		filepath.Join("conf", "conf-local.yaml"),
		filepath.Join("apps", "user", "conf", "conf-local.yaml"),
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
