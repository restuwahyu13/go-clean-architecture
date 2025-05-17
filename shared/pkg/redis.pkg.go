package pkg

import (
	"context"

	"time"

	goredis "github.com/redis/go-redis/v9"
	inf "github.com/restuwahyu13/go-clean-architecture/shared/interfaces"
)

type redis struct {
	redis *goredis.Client
	ctx   context.Context
}

func NewRedis(ctx context.Context, con *goredis.Client) (inf.IRedis, error) {
	return &redis{redis: con, ctx: ctx}, nil
}

func (p redis) SetEx(key string, expiration time.Duration, value any) error {
	cmd := p.redis.SetEx(p.ctx, key, value, expiration)

	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (p redis) Get(key string) ([]byte, error) {
	cmd := p.redis.Get(p.ctx, key)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	res := cmd.Val()
	return []byte(res), nil
}

func (p redis) Del(key string) (int64, error) {
	cmd := p.redis.Del(p.ctx, key)

	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return cmd.Val(), nil
}

func (p redis) Exists(key string) (int64, error) {
	cmd := p.redis.Exists(p.ctx, key)

	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return cmd.Val(), nil
}

func (p redis) HSetEx(key string, expiration time.Duration, values ...any) error {
	cmd := p.redis.HSet(p.ctx, key, values)
	p.redis.Expire(p.ctx, key, expiration)

	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (p redis) HGet(key, field string) ([]byte, error) {
	cmd := p.redis.HGet(p.ctx, key, field)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	res := cmd.Val()
	return []byte(res), nil
}
