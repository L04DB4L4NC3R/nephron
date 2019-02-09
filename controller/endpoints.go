package controller

import (
	"net/http"
)

func Startup() {
	http.HandleFunc("/upload", Fetch{}.UploadHandler())
	http.HandleFunc("/train", Fetch{}.TrainHandler())
	http.HandleFunc("/search", Fetch{}.FuzzySearch())
	http.HandleFunc("/get-logs-json", Fetch{}.GetLogsJSON())

}
