package model

import (
	"context"
	"encoding/json"
	"fmt"
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
			strr := strings.Replace(elem.Body, "\n", " ", -1)
			// make request
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(`{"title" : "` + elem.Title + `", "body":"` + strr + `"}`),
				Refresh:    "true",
			}

			// exec request
			res, err := req.Do(context.Background(), esc)
			c <- err
			defer res.Body.Close()
			fmt.Println(res)
			// deserialize response onto a map
			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
				c <- fmt.Errorf("Error indexing document ID")
				return
			} else {

				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
					c <- fmt.Errorf("Error parsing response body")
					return
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
			c <- nil
			return
		}(i, elem)
	}

}

// for one single value
func AddESDataSingle(key string, value string, c chan error) {
	v := strings.Replace(value, "\n", " ", -1)
	fmt.Println(value)
	go func() {

		// make request
		req := esapi.IndexRequest{
			Index: "test",
			// DocumentID: strconv.Itoa(1),
			Body:    strings.NewReader(fmt.Sprintf(`{"title" : "%s", "body":"%s"}`, key, v)),
			Refresh: "true",
		}

		// exec request
		res, err := req.Do(context.Background(), esc)
		if err != nil {
			c <- err
			return
		}
		defer res.Body.Close()
		fmt.Println(res)
		// deserialize response onto a map
		if res.IsError() {
			// log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			c <- fmt.Errorf("Error indexing document ID")
			return
		} else {

			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
				c <- fmt.Errorf("Error parsing response body")
				return
			} else {
				// Print the response status and indexed document version.
				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}
		c <- nil
		return
	}()

}
