package main

import (
	"fmt"
	"os"
)

func main() {
	//mendapatkan argumen di terminal
	arg := os.Args
	fmt.Println(arg)

	//mendapatkan hostname
	hostaName, err := os.Hostname()

	if err == nil {
		fmt.Println(hostaName)
	} else {
		fmt.Println("Error :", err.Error())
	}
}
