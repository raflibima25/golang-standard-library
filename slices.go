package main

import (
	"fmt"
	"slices"
)

func main() {
	// memanipulasi data dari slice

	names := []string{"Bima", "Raka", "Adi", "Putri"}
	values := []int{50, 70, 10, 30, 50}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Dudu"))
	fmt.Println(slices.Index(names, "Putri"))
	fmt.Println(slices.Index(names, "Dudu"))

}
