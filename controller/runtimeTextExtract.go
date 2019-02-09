package controller

import (
	"encoding/json"
	"net/http"

	"github.com/angadsharma1016/nephron/model"
	"github.com/angadsharma1016/nephron/services"
)

func (f Fetch) TrainHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get all data
		c := make(chan []model.ESdata)
		go services.All2text("data", c)
		msg := <-c
		close(c)

		// put in elasticsearch
		ce := make(chan error)
		go model.AddESdata(msg, ce)
		err := <-ce
		Must(err)

		// write to a file for testing purposes
		//err = ioutil.WriteFile("testing.txt", []byte(msg.Rs), 0644)
		//Must(err)

		//w.Write(msg.Rs)
		json.NewEncoder(w).Encode(struct {
			Status bool `json:"status"`
		}{true})

	}
}
