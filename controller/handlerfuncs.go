package controller

import (
	"net/http"
)

type Fetch struct {
	Rs string `json:"rs"`
}

func (f Fetch) UploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working"))
	}
}
