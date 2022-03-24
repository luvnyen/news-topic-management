package cache

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	host string
	db int
	expires time.Duration
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}