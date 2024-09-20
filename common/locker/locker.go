package locker

import (
	"context"
	"errors"
	"time"

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
}

func NewRedisLocker(ctx context.Context, client *redis.Client, key string) *RedisLocker {
	return &RedisLocker{ctx: ctx, client: client, key: key}
}

func (rl *RedisLocker) Lock() error {
	lockKey := lockPrefix + rl.key

	success, err := rl.client.SetNX(rl.ctx, lockKey, "locked", lockExpiry).Result()
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
	_, err := rl.client.Del(rl.ctx, lockKey).Result()

	return err
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
