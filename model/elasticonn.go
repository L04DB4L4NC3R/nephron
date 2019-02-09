package model

import (
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch"
)

var esc *elasticsearch.Client

func ConnectElastic() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	_, err = es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println("Connected to elasticsearch")

	esc = es
	return
}

func Must(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
