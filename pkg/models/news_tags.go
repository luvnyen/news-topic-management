package models

import (
	"github.com/jinzhu/gorm"
	"github.com/luvnyen/news-topic-management/pkg/config"
)

type NewsTags struct {
	gorm.Model
	NewsId int64 `json:"news_id"`
	TagId int64 `json:"tag_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&NewsTags{})
}

func (n *NewsTags) CreateNewsTags() *NewsTags {
	db.Create(n)
	return n
}

func GetAllNewsTags() []NewsTags {
	var newsTags []NewsTags
	db.Find(&newsTags)
	return newsTags
}

func GetTagsByNewsId(newsId int64) []Tags {
	var tag []Tags
	db.Table("news_tags").Select("tags.id, tags.name").Where("news_id = ?", newsId).Joins("JOIN tags ON tags.id = news_tags.tag_id").Find(&tag)
	return tag
}

func GetNewsTagsByNewsIdAndTagId(newsId int64, tagId int64) NewsTags {
	var newsTags NewsTags
	db.Where("news_id = ? AND tag_id = ?", newsId, tagId).First(&newsTags)
	return newsTags
}