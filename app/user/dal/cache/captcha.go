package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/aiagt/aiagt/kitex_gen/usersvc"
	"github.com/redis/go-redis/v9"
	"time"

	ktrdb "github.com/aiagt/kitextool/option/server/redis"
)

const (
	CaptchaKey = "captcha"
)

type CaptchaCache struct{}

func NewCaptchaCache() *CaptchaCache {
	return &CaptchaCache{}
}

func (c *CaptchaCache) rdb() *redis.Client {
	return ktrdb.RDB()
}

func (c *CaptchaCache) Set(ctx context.Context, typ CaptchaType, email, captcha string) error {
	key := fmt.Sprintf("%s:%s:%s", CaptchaKey, typ, email)
	return c.rdb().Set(ctx, key, captcha, 5*time.Minute).Err()
}

func (c *CaptchaCache) Get(ctx context.Context, typ CaptchaType, email string) (string, error) {
	key := fmt.Sprintf("%s:%s:%s", CaptchaKey, typ, email)

	captcha, err := c.rdb().Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", err
	}

	return captcha, nil
}

func (c *CaptchaCache) Del(ctx context.Context, typ CaptchaType, email string) error {
	key := fmt.Sprintf("%s:%s:%s", CaptchaKey, typ, email)
	return c.rdb().Del(ctx, key).Err()
}

func (c *CaptchaCache) GetAndDel(ctx context.Context, typ CaptchaType, email string) (string, error) {
	result, err := c.Get(ctx, typ, email)
	if err != nil {
		return "", err
	}

	_ = c.Del(ctx, typ, email)

	return result, nil
}

type CaptchaType string

var (
	CaptchaTypeAuth  = CaptchaType(usersvc.CaptchaType_AUTH.String())
	CaptchaTypeReset = CaptchaType(usersvc.CaptchaType_RESET.String())
)

func NewCaptchaType(t usersvc.CaptchaType) CaptchaType {
	return CaptchaType(t.String())
}
