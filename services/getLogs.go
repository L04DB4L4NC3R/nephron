package services

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/angadsharma1016/nephron/model"
)

func GetLogsJSON() []model.Logs {
	csvFile, _ := os.Open("logs.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []model.Logs
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, model.Logs{
			line[0],
			line[1],
			line[2],
			line[3],
			line[4],
		})
	}
	return people
}
