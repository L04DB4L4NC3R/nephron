package controller

import (
	"net/http"

	"github.com/angadsharma1016/nephron/model"
	"github.com/angadsharma1016/nephron/services"
)

func (f Fetch) TrainHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get all data
		c := make(chan model.ByteReturn)
		go services.All2text("data", c)
		msg := <-c
		Must(msg.Err)
		close(c)

		// put in elasticsearch
		ce := make(chan error)
		go model.AddESdata([]model.ESdata{model.ESdata{"hagga", "baba"}}, ce)
		err := <-ce
		Must(err)

		w.Write(msg.Rs)
		// json.NewEncoder(w).Encode(struct {
		// 	Status bool `json:"status"`
		// }{true})

	}
}
