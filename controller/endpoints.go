package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Startup() {
	router := mux.NewRouter()

	router.HandleFunc("/fetch", FetchFuzzySearch()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/upload", UploadHandler()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/train", TrainHandler()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/search", FuzzySearch()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/get-logs-json", GetLogsJSON()).Methods("GET", "OPTIONS", "POST")
	router.HandleFunc("/charts", ServeCharts()).Methods("GET", "OPTIONS", "POST")

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)

}
