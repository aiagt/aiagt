package confutil

import (
	ktconf "github.com/aiagt/kitextool/conf"
	"os"
	"path/filepath"
)

func LoadConf(conf ktconf.Conf, dirs ...string) {
	const (
		confFile        = "conf.yaml"
		confLocalFile   = "conf-local.yaml"
		confReleaseFile = "conf-release.yaml"
	)

	var confFiles []string

	for _, dir := range dirs {
		if IsReleaseEnv() {
			confFiles = append(confFiles, filepath.Join(dir, confReleaseFile))
		} else {
			confFiles = append(confFiles, filepath.Join(dir, confFile))
			confFiles = append(confFiles, filepath.Join(dir, confLocalFile))
		}
	}

	ktconf.LoadFiles(conf, confFiles...)
}

func IsReleaseEnv() bool {
	e := os.Getenv("GO_ENV")
	return e == "release"
}
