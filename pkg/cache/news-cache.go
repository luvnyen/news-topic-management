package cache

import (
	"encoding/json"
	"time"

	"github.com/luvnyen/news-topic-management/pkg/models"
)

type NewsCache interface {
	SetNews(key string, value *models.News)
	GetNews(key string) *models.News
}

func NewNewsRedisCache(host string, db int, expires time.Duration) NewsCache {
	return &redisCache{
		host: host,
		db: db,
		expires: expires,
	}
}

func (cache *redisCache) SetNews(key string, value *models.News) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires * time.Second)
}

func (cache *redisCache) GetNews(key string) *models.News {
	client := cache.getClient()

	data, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	news := &models.News{}
	err = json.Unmarshal([]byte(data), news)
	if err != nil {
		panic(err)
	}

	return news
}