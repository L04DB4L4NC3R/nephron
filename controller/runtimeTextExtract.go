package controller

import (
	"net/http"

	"github.com/angadsharma1016/nephron/model"
	"github.com/angadsharma1016/nephron/services"
)

func (f Fetch) TrainHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c := make(chan model.ByteReturn)
		go services.All2text("data", c)
		msg := <-c
		Must(msg.Err)
		close(c)
		w.Write(msg.Rs)

	}
}
