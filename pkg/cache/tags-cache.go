package cache

import (
	"encoding/json"
	"time"

	"github.com/luvnyen/news-topic-management/pkg/models"
)

type TagsCache interface {
	SetTags(key string, value *models.Tags)
	GetTags(key string) *models.Tags
}

func NewTagsRedisCache(host string, db int, expires time.Duration) TagsCache {
	return &redisCache{
		host: host,
		db: db,
		expires: expires,
	}
}

func (cache *redisCache) SetTags(key string, value *models.Tags) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires * time.Second)
}

func (cache *redisCache) GetTags(key string) *models.Tags {
	client := cache.getClient()

	data, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	tags := &models.Tags{}
	err = json.Unmarshal([]byte(data), tags)
	if err != nil {
		panic(err)
	}

	return tags
}