package email

import (
	"path/filepath"
	"testing"

	"github.com/aiagt/aiagt/app/user/conf"
	"github.com/aiagt/aiagt/app/user/pkg/captcha"
	ktconf "github.com/aiagt/kitextool/conf"
	"github.com/stretchr/testify/assert"
)

func TestSendAuthCaptcha(t *testing.T) {
	ktconf.LoadFiles(conf.Conf(), filepath.Join("..", "..", "conf", "conf.yaml"))

	err := SendAuthCaptcha(captcha.Generate(), "ahao_study@163.com")
	assert.NoError(t, err)
}
