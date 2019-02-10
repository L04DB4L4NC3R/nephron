package controller

import (
	"github.com/gorilla/mux"
)

func Startup() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/upload", Fetch{}.UploadHandler()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/train", Fetch{}.TrainHandler()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/search", Fetch{}.FuzzySearch()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/get-logs-json", Fetch{}.GetLogsJSON()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/search-excel", Fetch{}.ExcelFuzzySearch()).Methods("GET", "OPTIONS", "POST")
	return router

}
