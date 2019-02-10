package model

import (
	"fmt"
	"log"
	"os"
)

func RecordLog(logstring LogType) {
	f, err := os.OpenFile("logs.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(fmt.Sprintf(",%s,%d,%f,%f", logstring.Query, logstring.Hits, logstring.Time, logstring.MaxScore))

}
