package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		fmt.Println(msg.Rs[0]["_source"])
		w.Write([]byte("DONE"))
	}
}
