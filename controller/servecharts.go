package controller

import (
	"fmt"
	"log"
	"net/http"
)

func ServeCharts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm() // ParseMultipartForm to deal with multipart data, it calls ParseForm also
		if err != nil {
			fmt.Println(err)
		}
		f := r.Form
		path := "static/" + f.Get("query") + ".html"
		log.Println(path)
		// http.ServeFile(w, r, path)
	}
}
