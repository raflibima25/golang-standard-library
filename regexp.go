package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`bi([a-z])a`)

	fmt.Println(regex.MatchString("bima"))
	fmt.Println(regex.MatchString("bibima"))
	fmt.Println(regex.MatchString("bImBima"))

	fmt.Println(regex.FindAllString("bima bibima bisa bija bika bi2a", 10))
}
