package locker

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const (
	lockPrefix = "distributed_lock:"
	lockExpiry = 3 * time.Second
)

type RedisLocker struct {
	ctx    context.Context
	client *redis.Client
	key    string
	value  string
}

func NewRedisLocker(ctx context.Context, client *redis.Client, key string) *RedisLocker {
	return &RedisLocker{
		ctx:    ctx,
		client: client,
		key:    key,
		value:  uuid.NewString(),
	}
}

func (rl *RedisLocker) Lock() error {
	lockKey := lockPrefix + rl.key

	success, err := rl.client.SetNX(rl.ctx, lockKey, rl.value, lockExpiry).Result()
	if err != nil {
		return err
	}

	if !success {
		return errors.New("failed to acquire lock")
	}

	return nil
}

func (rl *RedisLocker) Unlock() error {
	lockKey := lockPrefix + rl.key

	luaScript := `
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`

	res, err := rl.client.Eval(rl.ctx, luaScript, []string{lockKey}, rl.value).Int()
	if err != nil {
		return err
	}

	if res == 0 {
		return errors.New("failed to release lock or lock does not exist")
	}

	return nil
}

func (rl *RedisLocker) LockWithRetry(retryInterval time.Duration, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := rl.Lock()
		if err == nil {
			return nil
		}
		select {
		case <-rl.ctx.Done():
			return rl.ctx.Err()
		case <-time.After(retryInterval):
		}
	}

	return errors.New("max retries reached, failed to acquire lock")
}
