package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/aiagt/aiagt/common/locker"
	"github.com/aiagt/aiagt/pkg/logerr"
	ktrdb "github.com/aiagt/kitextool/option/server/redis"
)

const (
	CallTokenKey = "call_token"
)

type CallTokenCache struct{}

func NewCallTokenCache() *CallTokenCache {
	return &CallTokenCache{}
}

func (c *CallTokenCache) rdb() *redis.Client {
	return ktrdb.RDB()
}

type CallTokenValue struct {
	AppID          int64  `json:"app_id"`
	PluginID       *int64 `json:"plugin_id,omitempty"`
	ConversationID int64  `json:"conversation_id"`
	CallLimit      int32  `json:"call_limit"`
}

func (v *CallTokenValue) MarshalBinary() ([]byte, error) {
	return json.Marshal(v)
}

func (v *CallTokenValue) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, v)
}

func (c *CallTokenCache) Set(ctx context.Context, token string, val *CallTokenValue) error {
	key := fmt.Sprintf("%s:%s", CallTokenKey, token)
	return c.rdb().Set(ctx, key, val, 10*time.Minute).Err()
}

func (c *CallTokenCache) Get(ctx context.Context, token string) (*CallTokenValue, error) {
	key := fmt.Sprintf("%s:%s", CallTokenKey, token)

	result := c.rdb().Get(ctx, key)
	if err := result.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	var val CallTokenValue

	err := result.Scan(&val)
	if err != nil {
		return nil, err
	}

	return &val, nil
}

func (c *CallTokenCache) Decr(ctx context.Context, token string) (bool, error) {
	var (
		key  = fmt.Sprintf("%s:%s", CallTokenKey, token)
		lock = locker.NewRedisLocker(ctx, c.rdb(), key)
	)

	err := lock.LockWithRetry(100*time.Millisecond, 10)
	if err != nil {
		return false, err
	}
	defer logerr.Log(lock.Unlock())

	val, err := c.Get(ctx, token)
	if err != nil {
		return false, err
	}

	if val.CallLimit <= 0 {
		return false, nil
	}

	val.CallLimit--

	return true, c.Set(ctx, token, val)
}
