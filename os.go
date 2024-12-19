package main

import (
	"fmt"
	"os"
)

func main() {
	// args
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}
	// go run os.go rafli bima pratandra
	// go run os.go "rafli bima pratandra"
	// di belakang os.go adalah argumen

	// hostname
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

}
