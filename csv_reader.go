package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "rafli,bima,pratandra\n" +
		"bimbi,kamal,riyadi\n" +
		"raka,trisatya,nugraha"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
