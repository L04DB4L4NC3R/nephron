package controller

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/angadsharma1016/nephron/model"
	"github.com/angadsharma1016/nephron/services"
)

type Fetch struct {
	Rs string `json:"rs"`
}

func (f Fetch) UploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
		fl, err := os.OpenFile("data/"+name, os.O_RDWR|os.O_CREATE, 0755)
		Must(err)
		defer fl.Close()
		io.Copy(fl, file)

		// if name is excel then save filename as key and filepath as value
		ext := strings.Split(name, ".")[1]
		if ext == "xls" || ext == "xlsx" || ext == "csv" {
			c := make(chan error)
			go model.AddESDataSingle(name, "data/"+name, c)
			err := <-c
			close(c)
			Must(err)
			w.Write([]byte("DONE"))
			return
		}

		// goroutine for converting to text
		c := make(chan model.StringReturn)
		go services.ConvertToText("data/"+name, c)
		msg := <-c
		close(c)
		Must(msg.Err)

		cr := make(chan error)
		model.AddESDataSingle(name, msg.Rs, cr)
		Must(<-cr)
		w.Write([]byte(msg.Rs))

		return

	}
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
