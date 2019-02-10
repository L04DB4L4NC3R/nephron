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

		http.ServeFile(w, r, path)
		http.Handle("/images/", http.FileServer(http.Dir("static")))
		http.Handle("/stylesheets/", http.FileServer(http.Dir("static")))
		http.Handle("/js/", http.FileServer(http.Dir("static")))
	}
}
