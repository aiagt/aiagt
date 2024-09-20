package captcha

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Generate() string {
	return fmt.Sprintf("%6d", r.Intn(999999))
}
