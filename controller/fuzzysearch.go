package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/angadsharma1016/nephron/model"
)

func (f Fetch) FuzzySearch() http.HandlerFunc {
	// this link only if no xls, xlsx, csv
	return func(w http.ResponseWriter, r *http.Request) {
		var q model.Query
		json.NewDecoder(r.Body).Decode(&q)
		c := make(chan model.FuzzyReturn)
		go model.FuzzySearch(q.Query, c)

		msg := <-c
		Must(msg.Err)
		if len(msg.Rs) == 0 {
			log.Println("Search miss")
			w.Write([]byte("Search miss"))
			return
		}

		w.Write([]byte(msg.Rs[0]["_source"].(map[string]interface{})["body"].(string)))

	}
}

// handles excel files and forms
func (f Fetch) ExcelFuzzySearch() http.HandlerFunc {
	// this link only if no xls, xlsx, csv
	return func(w http.ResponseWriter, r *http.Request) {

		err := r.ParseForm() // ParseMultipartForm to deal with multipart data, it calls ParseForm also
		if err != nil {
			fmt.Println(err)
		}
		f := r.Form
		c := make(chan model.FuzzyReturn)

		go model.FuzzySearch(f.Get("query"), c)

		msg := <-c
		Must(msg.Err)
		if len(msg.Rs) == 0 {
			log.Println("Search miss")
			w.Write([]byte("Search miss"))
			return
		}

		http.ServeFile(w, r, msg.Rs[0]["_source"].(map[string]interface{})["body"].(string))

		//w.Write([]byte(msg.Rs[0]["_source"].(map[string]interface{})["body"].(string)))

	}
}
