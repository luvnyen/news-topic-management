package controllers

import (
	"testing"

	"github.com/luvnyen/news-topic-management/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetTagsById(t *testing.T) {
	tag := models.Tags{
		Name: "Test",
	}
	tag.CreateTags()

	tags_by_id, _ := models.GetTagsById(int64(tag.ID))

	assert.Equal(t, tag.ID, tags_by_id.ID, "Data not found")

	models.DeleteTags(int64(tag.ID))
}

func TestCreateTags(t *testing.T) {
	curr_tags_len := len(models.GetAllTags())

	tag := models.Tags{
		Name: "Test",
	}
	tag.CreateTags()
	tags_by_id, _ := models.GetTagsById(int64(tag.ID))

	assert.Equal(t, tag.ID, tags_by_id.ID, "Data not found")
	assert.Equal(t, curr_tags_len+1, len(models.GetAllTags()), "Length not increased")

	models.DeleteTags(int64(tag.ID))
}

func TestDeleteTags(t *testing.T) {
	tag := models.Tags{
		Name: "Test",
	}
	tag.CreateTags()

	curr_tags_len := len(models.GetAllTags())
	models.DeleteTags(int64(tag.ID))
	tags_by_id, _ := models.GetTagsById(int64(tag.ID))

	assert.Equal(t, uint(0), tags_by_id.ID, "Data not deleted")
	assert.Equal(t, curr_tags_len-1, len(models.GetAllTags()), "Length not decreased")
}