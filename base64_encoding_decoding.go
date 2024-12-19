package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// jika ingin file
	filePath := "CV Winah Aulia.pdf"
	_, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err.Error())
	}

	value := "Rafli Bima Pratandra"

	var encoded = base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatalf("failed to decode %v", err.Error())
	}
	fmt.Println(string(decoded))
}
