package controllers

import (
	"testing"

	"github.com/luvnyen/news-topic-management/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetNewsById(t *testing.T) {
	news := models.News{
		Title: "Test",
		Status: "publish",
	}
	news.CreateNews()
	
	news_by_id, _ := models.GetNewsById(int64(news.ID))
	
	assert.Equal(t, news.ID, news_by_id.ID, "Data not found")

	models.DeleteNews(int64(news.ID))
}

func TestGetNewsByStatus(t *testing.T) {
	news := models.News{
		Title: "Test",
		Status: "draft",
	}
	news.CreateNews()

	news_by_status := models.GetNewsByStatus("draft")

	assert.Equal(t, news.Status, news_by_status[0].Status, "Data not found")

	models.DeleteNews(int64(news.ID))
}

func TestGetNewsByTopic(t *testing.T) {
	news := models.News{
		Title: "Investment in the future",
		Status: "publish",
	}
	news.CreateNews()

	topic := "investment"
	news_by_topic := models.GetNewsByTopic(topic)

	assert.Contains(t, news_by_topic[0].Title, topic, "Data not found")

	models.DeleteNews(int64(news.ID))
}

func TestCreateNews(t *testing.T) {
	curr_news_len := len(models.GetAllNews())

	news := models.News{
		Title: "Test",
		Status: "Test",
	}
	news.CreateNews()
	news_by_id, _ := models.GetNewsById(int64(news.ID))

	assert.Equal(t, news.ID, news_by_id.ID, "Data not found")
	assert.Equal(t, curr_news_len+1, len(models.GetAllNews()), "Length not increased")

	models.DeleteNews(int64(news.ID))
}

func TestDeleteNews(t *testing.T) {
	news := models.News{
		Title: "Test",
		Status: "Test",
	}
	news.CreateNews()

	curr_news_len := len(models.GetAllNews())
	models.DeleteNews(int64(news.ID))
	news_by_id, _ := models.GetNewsById(int64(news.ID))
	
	assert.Equal(t, uint(0), news_by_id.ID, "Data not deleted")
	assert.Equal(t, curr_news_len-1, len(models.GetAllNews()), "Length not decreased")
}