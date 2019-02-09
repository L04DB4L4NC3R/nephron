package model

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/esapi"
)

var (
	r  map[string]interface{}
	wg sync.WaitGroup
)

func AddESdata(data []ESdata, c chan error) {

	for i, elem := range data {

		// put req in waiting group
		wg.Add(1)

		go func(i int, elem ESdata) {
			defer wg.Done()

			// make request
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(`{"title" : "` + elem.Title + `", "body":"` + elem.Body + `"}`),
				Refresh:    "true",
			}

			// exec request
			res, err := req.Do(context.Background(), esc)
			Must(err)
			defer res.Body.Close()

			// deserialize response onto a map
			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, elem)
	}

}
