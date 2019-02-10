package main

import (
	"github.com/angadsharma1016/nephron/controller"
	"github.com/angadsharma1016/nephron/model"
)

func main() {
	model.ConnectElastic()
	controller.Startup()

}
