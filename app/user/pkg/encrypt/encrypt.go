package encrypt

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/aiagt/aiagt/app/user/conf"
)

func Encrypt(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + conf.Conf().Auth.EncryptSalt))

	return hex.EncodeToString(hash.Sum(nil))
}
