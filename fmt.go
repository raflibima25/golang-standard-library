package main

import "fmt"

func main() {
	firstName := "Bima"
	lastName := "Pratandra"

	fmt.Println("Hello '", firstName, lastName, "'") // pada kutip terdapat spasi

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)

}
