package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Startup() {
	router := mux.NewRouter()

	router.HandleFunc("/upload", Fetch{}.UploadHandler()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/train", Fetch{}.TrainHandler()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/search", Fetch{}.FuzzySearch()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/get-logs-json", Fetch{}.GetLogsJSON()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/search-excel", Fetch{}.ExcelFuzzySearch()).Methods("GET", "OPTIONS", "POST")

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)

}
