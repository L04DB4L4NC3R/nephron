package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func FuzzySearch(query string, c chan FuzzyReturn) {

	// search with fuzziness 2
	res, err := esc.Search(
		esc.Search.WithContext(context.Background()),
		esc.Search.WithIndex("test"),
		esc.Search.WithBody(strings.NewReader(`{"query":{"match":{"body":{"query":"`+query+`", "fuzziness":"2"}}}}`)),
		esc.Search.WithTrackTotalHits(true),
		esc.Search.WithPretty(),
	)
	Must(err)
	defer res.Body.Close()

	// unmarshal search body
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			c <- FuzzyReturn{[]map[string]interface{}{}, err}
			log.Printf("error parsing the response body: %s", err)
			return
		} else {
			// Print the response status and error information.
			c <- FuzzyReturn{nil, fmt.Errorf("Error searching")}

			log.Printf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	var arr []map[string]interface{}
	if r["hits"] == nil {
		c <- FuzzyReturn{arr, fmt.Errorf("No result found")}
		return
	}

	// logdata- hits and time taken
	if r["hits"].(map[string]interface{})["total"].(float64) != 0 {
		RecordLog(LogType{
			query,
			int(r["hits"].(map[string]interface{})["total"].(float64)),
			r["took"].(float64),
			r["hits"].(map[string]interface{})["max_score"].(float64),
		})

	} else {
		RecordLog(LogType{
			query,
			0,
			0,
			0,
		})
		c <- FuzzyReturn{arr, nil}
		return
	}

	// collect hits
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

		arr = append(arr, hit.(map[string]interface{}))
		//log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
	c <- FuzzyReturn{arr, nil}
	return

}
