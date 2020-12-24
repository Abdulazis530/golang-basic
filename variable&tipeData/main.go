package main

import "fmt"

func endApp() {
	fmt.Println("endApp")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERRRROR")
	}
}

func main() {
	// runApp(true)

	type Customer struct {
		Name, Adress string
		Age          int
	}

	customer1 := Customer{"AZIS", "PRIOK", 25}
	customer2 := customer1

	customer2.Adress = "AMERIKA"
	fmt.Println(customer1)
	fmt.Println(customer2)
}
