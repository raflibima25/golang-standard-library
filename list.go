package main

import (
	"container/list"
	"fmt"
)

// Struktur data Double Linked List

func main() {
	var data *list.List = list.New()

	data.PushBack("rafli")
	data.PushBack("bima")
	data.PushBack("pratandra")
	data.PushFront("bibih")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // rafli

	next := head.Next()
	fmt.Println(next.Value) // bima

	element := next.Next()
	fmt.Println(element.Value) // pratandra

	// cara menggunakan iterasi
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println("iterasi double linked lis:", e.Value)
	}

}
