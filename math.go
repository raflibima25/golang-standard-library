package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Floor(1.40))
	fmt.Println(math.Ceil(1.60))
	fmt.Println(math.Round(1.60))
	fmt.Println(math.Min(11, 12))
	fmt.Println(math.Max(11, 12))
}
