package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luvnyen/news-topic-management/pkg/cache"
	"github.com/luvnyen/news-topic-management/pkg/models"
)

// stored in redis for 120s
var newsCache cache.NewsCache = cache.NewNewsRedisCache("localhost:6379", 1, 120)

func GetNews(w http.ResponseWriter, r *http.Request) {
	news := models.GetAllNews()
	if len(news) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("News not found!"))
		return
	}
	
	res, _ := json.Marshal(news)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetNewsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	newsId := params["newsId"]
	ID, err := strconv.ParseInt(newsId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}
	
	// check if exist in cache, else get from database
	news := newsCache.GetNews(newsId)
	if news.ID == 0 {
		news, _ := models.GetNewsById(ID)
		if news.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("News not found!"))
			return
		}
		newsCache.SetNews(newsId, news)

		res, _ := json.Marshal(news)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		res, _ := json.Marshal(news)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetNewsByStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status := params["newsStatus"]
	if status != "publish" && status != "draft" && status != "deleted" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid status!"))
		return
	}

	news := models.GetNewsByStatus(status)
	if len(news) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("News not found!"))
		return
	}
	
	res, _ := json.Marshal(news)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetNewsByTopic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	newsTopic := params["newsTopic"]

	news := models.GetNewsByTopic(newsTopic)
	if len(news) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("News not found!"))
		return
	}

	res, _ := json.Marshal(news)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNews(w http.ResponseWriter, r *http.Request) {
	news := &models.News{}
	err := json.NewDecoder(r.Body).Decode(news)
	if err != nil {
		fmt.Println("Error while decoding!")
	}
	if news.Title == "" || news.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data!"))
		return
	}
	if news.Status != "publish" && news.Status != "draft" && news.Status != "deleted" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid status!"))
		return
	}
	n := news.CreateNews()

	res, _ := json.Marshal(n)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	news := &models.News{}
	err := json.NewDecoder(r.Body).Decode(news)
	if err != nil {
		fmt.Println("Error while decoding!")
	}
	if news.Title == "" || news.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data!"))
		return
	}
	if news.Status != "publish" && news.Status != "draft" && news.Status != "deleted" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid status!"))
		return
	}
	
	params := mux.Vars(r)
	newsId := params["newsId"]
	ID, err := strconv.ParseInt(newsId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}

	newsDetails, db := models.GetNewsById(ID)
	if newsDetails.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("News not found!"))
		return
	}
	if news.Title != "" {
		newsDetails.Title = news.Title
	}
	if news.Status != "" {
		newsDetails.Status = news.Status
	}
	db.Save(newsDetails)

	res, _ := json.Marshal(newsDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteNews(w http.ResponseWriter, r *http.Request) {
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
	n := models.DeleteNews(ID)

	res, _ := json.Marshal(n)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}