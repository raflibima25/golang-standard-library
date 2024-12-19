package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/main.go"))
	fmt.Println(path.Base("hello/main.go"))
	fmt.Println(path.Ext("hello/main.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))

}
