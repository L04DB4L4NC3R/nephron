package services

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/angadsharma1016/nephron/model"
)

func All2text(dir string, cr chan model.ByteReturn) {
	files, err := ioutil.ReadDir(dir)
	Must(err)

	c := make(chan model.StringReturn, len(files))

	// convert all files in the dir to text
	for _, f := range files {
		log.Println(f.Name())
		go ConvertToText(dir+"/"+f.Name(), c)
	}

	var buf bytes.Buffer
	count := 0

	// range over channel buffers and append them
	// then flush once done
	for elem := range c {
		if elem.Err != nil {
			close(c)
			cr <- model.ByteReturn{nil, elem.Err}
			return
		}
		count++
		buf.WriteString(elem.Rs)

		if count == len(files) {
			close(c)
			cr <- model.ByteReturn{buf.Bytes(), nil}
			buf.Reset()
			return
		}
	}

}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
