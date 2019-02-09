package controller

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/angadsharma1016/nephron/model"
	"github.com/angadsharma1016/nephron/services"
)

type Fetch struct {
	Rs string `json:"rs"`
}

func (f Fetch) UploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.Write([]byte("Method not allowed"))
			return
		}

		// accept file
		file, header, err := r.FormFile("file")
		Must(err)
		defer file.Close()
		name := header.Filename

		// create a new file and pipe to it
		fl, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
		Must(err)
		defer fl.Close()
		io.Copy(fl, file)

		// goroutine for converting to text
		c := make(chan model.StringReturn)
		go services.ConvertToText(name, c)
		msg := <-c
		Must(msg.Err)
		w.Write([]byte(msg.Rs))

		return

	}
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
