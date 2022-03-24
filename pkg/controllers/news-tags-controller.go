package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luvnyen/news-topic-management/pkg/models"
)

func GetNewsTags(w http.ResponseWriter, r *http.Request) {
	newsTags := models.GetAllNewsTags()
	if len(newsTags) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("News tags not found!"))
		return
	}

	res, _ := json.Marshal(newsTags)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTagsByNewsId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	newsId := params["newsId"]
	ID, err := strconv.ParseInt(newsId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}

	news, _ := models.GetNewsById(ID)
	if news.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("News not found!"))
		return
	}

	tags := models.GetTagsByNewsId(ID)
	if len(tags) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tags not found!"))
		return
	}

	res, _ := json.Marshal(tags)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNewsTags(w http.ResponseWriter, r *http.Request) {
	var newsTags models.NewsTags
	json.NewDecoder(r.Body).Decode(&newsTags)

	if newsTags.NewsId == 0 || newsTags.TagId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data!"))
		return
	}

	news, _ := models.GetNewsById(newsTags.NewsId)
	if news.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("News not found!"))
		return
	}
	tags, _ := models.GetTagsById(newsTags.TagId)
	if tags.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tag not found!"))
		return
	}

	newsTags = models.GetNewsTagsByNewsIdAndTagId(newsTags.NewsId, newsTags.TagId)
	if newsTags.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("News and tag already exist!"))
		return
	}

	newsTags.CreateNewsTags()

	res, _ := json.Marshal(newsTags)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}