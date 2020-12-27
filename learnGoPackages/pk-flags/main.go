package main

import (
	"flag"
	"fmt"
)

//package flags berisikan fungsionalitas untuk memparsing command line argumnets
func main() {
	var host *string = flag.String("host", "localhost", "put your host")
	var user *string = flag.String("user", "localhost", "put your database user")
	var password *string = flag.String("password", "localhost", "put your database password")

	flag.Parse()
	fmt.Println("Host: *", *host)
	fmt.Println("User: *", *user)
	fmt.Println("Password: *", *password)
}
