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

var tagsCache cache.TagsCache = cache.NewTagsRedisCache("localhost:6379", 2, 120)

func GetTags(w http.ResponseWriter, r *http.Request) {
	tags := models.GetAllTags()
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

func GetTagsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tagsId := params["tagsId"]
	ID, err := strconv.ParseInt(tagsId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}

	// check if exist in cache, else get from database
	tags := tagsCache.GetTags(tagsId)
	if tags.ID == 0 {
		tags, _ := models.GetTagsById(ID)
		if tags.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Tags not found!"))
			return
		}
		tagsCache.SetTags(tagsId, tags)

		res, _ := json.Marshal(tags)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		res, _ := json.Marshal(tags)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateTags(w http.ResponseWriter, r *http.Request) {
	tags := &models.Tags{}
	err := json.NewDecoder(r.Body).Decode(tags)
	if err != nil {
		fmt.Println("Error while decoding!")
	}
	if tags.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data!"))
		return
	}
	_, db := models.GetTagsByName(tags.Name)
	if !db.NewRecord(tags) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Tags name already exist!"))
		return
	}

	n := tags.CreateTags()

	res, _ := json.Marshal(n)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateTags(w http.ResponseWriter, r *http.Request) {
	tags := &models.Tags{}
	err := json.NewDecoder(r.Body).Decode(tags)
	if err != nil {
		fmt.Println("Error while decoding!")
	}
	if tags.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data!"))
		return
	}
	
	params := mux.Vars(r)
	tagsId := params["tagsId"]
	ID, err := strconv.ParseInt(tagsId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}

	tagsDetails, db := models.GetTagsById(ID)
	if tagsDetails.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tags not found!"))
		return
	}
	if tags.Name != "" {
		tagsDetails.Name = tags.Name
	}
	db.Save(tagsDetails)

	res, _ := json.Marshal(tagsDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTags(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tagsId := params["tagsId"]
	ID, err := strconv.ParseInt(tagsId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}

	tags, _ := models.GetTagsById(ID)
	if tags.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Tags not found!"))
		return
	}
	n := models.DeleteTags(ID)

	res, _ := json.Marshal(n)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}