package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Rafli Bima", "Bima"))
	fmt.Println(strings.ToLower("Rafli Bima"))
	fmt.Println(strings.ToUpper("Rafli Bima"))
	fmt.Println(strings.Split("Rafli Bima", " "))
	fmt.Println(strings.Trim("               Rafli Bima              ", " "))
	fmt.Println(strings.ReplaceAll("Rafli Bima Bima Rafli", "Rafli", "Manis"))
}
