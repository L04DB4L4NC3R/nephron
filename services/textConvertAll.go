package services

import (
	"io/ioutil"
	"log"

	"github.com/angadsharma1016/nephron/model"
)

func All2text(dir string, cr chan []model.ESdata) {
	files, err := ioutil.ReadDir(dir)
	Must(err)

	c := make(chan model.StringReturn, len(files))

	var filenames []string
	// convert all files in the dir to text
	for _, f := range files {
		log.Println(f.Name())
		filenames = append(filenames, f.Name())
		go ConvertToText(dir+"/"+f.Name(), c)
	}

	count := 0
	var esarr []model.ESdata
	// range over channel buffers and append them
	// then flush once done
	for elem := range c {
		if elem.Err != nil {
			close(c)
			log.Fatal(elem.Err)
			cr <- esarr
			return
		}
		esarr = append(esarr, model.ESdata{filenames[count], elem.Rs})
		count++

		if count == len(files) {
			close(c)
			cr <- esarr
			return
		}
	}

}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
