package models

import (
	"github.com/jinzhu/gorm"
	"github.com/luvnyen/news-topic-management/pkg/config"
)

var db *gorm.DB

type News struct {
	gorm.Model
	Title string `json:"title"`
	Status string `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&News{})
}

func (n *News) CreateNews() *News {
	db.Create(n)
	return n
}

func GetAllNews() []News {
	var news []News
	db.Find(&news)
	return news
}

func GetNewsById(id int64) (*News, *gorm.DB) {
	var news News
	db := db.Where("id = ?", id).First(&news)
	return &news, db
}

func GetNewsByStatus(status string) []News {
	var news []News
	db.Where("status = ?", status).Find(&news)
	return news
}

func GetNewsByTopic(topic string) []News {
	var news []News
	db.Where("title LIKE ?", "%"+topic+"%").Find(&news)
	return news
}

func DeleteNews(id int64) News {
	var news News
	db.Where("id = ?", id).Delete(&news)
	return news
}