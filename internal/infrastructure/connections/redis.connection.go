package con

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/restuwahyu13/go-clean-architecture/shared/dto"
)

func RedisConnection(env dto.Environtment) (*redis.Client, error) {
	parseURL, err := redis.ParseURL(env.REDIS.URL)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(&redis.Options{
		Addr:            parseURL.Addr,
		MaxRetries:      10,
		PoolSize:        20,
		PoolFIFO:        true,
		ReadTimeout:     time.Duration(time.Second * 30),
		WriteTimeout:    time.Duration(time.Second * 30),
		DialTimeout:     time.Duration(time.Second * 60),
		MinRetryBackoff: time.Duration(time.Second * 60),
		MaxRetryBackoff: time.Duration(time.Second * 120),
	}), nil
}
