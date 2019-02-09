package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/angadsharma1016/nephron/model"
)

func (f Fetch) FuzzySearch() http.HandlerFunc {
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

		if ext := strings.Split(msg.Rs[0]["_source"].(map[string]interface{})["body"].(string), ".")[1]; ext == "xls" || ext == "xlsx" || ext == "csv" {
			http.ServeFile(w, r, msg.Rs[0]["_source"].(map[string]interface{})["body"].(string))
		} else {
			w.Write([]byte(msg.Rs[0]["_source"].(map[string]interface{})["body"].(string)))
		}

	}
}
