package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 5432, "database port")

	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)

	// go run flag.go -username=bima -password="rahasia banget" -host=10.31.42.6 -port=1234
}
