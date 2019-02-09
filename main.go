// package main

// import (
// 	"fmt"
// 	"log"

// 	"code.sajari.com/docconv/client"
// )

// func main() {
// 	c := client.New()
// 	res, err := client.ConvertPath(c, `test.pdf`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(res)
// }

package main

import (
	"log"
	"net/http"

	"github.com/angadsharma1016/nephron/controller"
	"github.com/angadsharma1016/nephron/model"
)

func main() {
	model.ConnectElastic()

	c := make(chan error)

	r := []model.ESdata{
		model.ESdata{"angad", "sharma"},
		model.ESdata{"dhruv", "sharma"},
		model.ESdata{"rakesh", "sharma"},
	}
	model.AddESdata(r, c)

	controller.Startup()
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
