package main

import (
	"log"
	"net/http"

	"github.com/angadsharma1016/nephron/controller"
	"github.com/angadsharma1016/nephron/model"
)

func main() {
	model.ConnectElastic()
	muxx := controller.Startup()
	log.Println("Listening...")
	http.ListenAndServe(":3000", muxx)
}
