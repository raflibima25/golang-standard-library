package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parseBool)
	}

	parseInt, err := strconv.Atoi("10000")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parseInt)
	}

	formatInt := strconv.FormatInt(999, 2)
	fmt.Println(formatInt)

	itoa := strconv.Itoa(999)
	fmt.Println(itoa)
}
