package routes

import (
	"github.com/gorilla/mux"
	"github.com/luvnyen/news-topic-management/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/news/", controllers.GetNews).Methods("GET")
	router.HandleFunc("/news/{newsId}", controllers.GetNewsById).Methods("GET")
	router.HandleFunc("/news/topic/{newsTopic}", controllers.GetNewsByTopic).Methods("GET")
	router.HandleFunc("/news/status/{newsStatus}", controllers.GetNewsByStatus).Methods("GET")
	router.HandleFunc("/news/", controllers.CreateNews).Methods("POST")
	router.HandleFunc("/news/{newsId}", controllers.UpdateNews).Methods("PUT")
	router.HandleFunc("/news/{newsId}", controllers.DeleteNews).Methods("DELETE")

	router.HandleFunc("/tags/", controllers.GetTags).Methods("GET")
	router.HandleFunc("/tags/{tagsId}", controllers.GetTagsById).Methods("GET")
	router.HandleFunc("/tags/", controllers.CreateTags).Methods("POST")
	router.HandleFunc("/tags/{tagsId}", controllers.UpdateTags).Methods("PUT")
	router.HandleFunc("/tags/{tagsId}", controllers.DeleteTags).Methods("DELETE")

	router.HandleFunc("/news/tags/", controllers.GetNewsTags).Methods("GET")
	router.HandleFunc("/news/tags/{newsId}", controllers.GetTagsByNewsId).Methods("GET")
	router.HandleFunc("/news/tags/", controllers.CreateNewsTags).Methods("POST")
}