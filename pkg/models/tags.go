package models

import (
	"github.com/jinzhu/gorm"
	"github.com/luvnyen/news-topic-management/pkg/config"
)

type Tags struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Tags{})
}

func (n *Tags) CreateTags() *Tags {
	db.Create(n)
	return n
}

func GetAllTags() []Tags {
	var Tags []Tags
	db.Find(&Tags)
	return Tags
}

func GetTagsById(id int64) (*Tags, *gorm.DB) {
	var Tags Tags
	db := db.Where("id = ?", id).First(&Tags)
	return &Tags, db
}

func GetTagsByName(name string) (*Tags, *gorm.DB) {
	var Tags Tags
	db := db.Where("name = ?", name).First(&Tags)
	return &Tags, db
}

func DeleteTags(id int64) Tags {
	var Tags Tags
	db.Where("id = ?", id).Delete(&Tags)
	return Tags
}