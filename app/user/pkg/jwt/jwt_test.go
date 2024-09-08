package jwt

import (
	"github.com/aiagt/aiagt/app/user/conf"
	ktconf "github.com/aiagt/kitextool/conf"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	ktconf.LoadFiles(conf.Conf(), filepath.Join("..", "..", "conf", "conf.yaml"))

	token, _, err := GenerateToken(128)
	assert.NoError(t, err)

	id, err := ParseToken(token)
	assert.NoError(t, err)

	assert.Equal(t, id, int64(128))
}
