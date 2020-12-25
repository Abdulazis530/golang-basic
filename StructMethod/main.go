package main

import "fmt"

//menambahkan method kedalam stuct

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
func main() {
	firaun := Customer{
		Name:    "Firaun",
		Address: "Mesir",
		Age:     22222,
	}

	firaun.sayHello("Hitler")
}
