package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("this is bufio\nwith new reader")
	bufReader := bufio.NewReader(reader)

	for {
		line, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		}

		//fmt.Println(prefix)
		fmt.Println(string(line))
	}
}
