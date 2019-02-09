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
	"net/http"

	"github.com/angadsharma1016/nephron/controller"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":3000", nil)
}
