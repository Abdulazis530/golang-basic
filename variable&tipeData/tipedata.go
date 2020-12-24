package main

import "fmt"

func main() {
	// var angka int = 123
	// var desimal float32 = 1.23
	// var bolecuria bool = true
	// fmt.Printf("angka = %d \n", angka)
	// fmt.Printf("angka decimal= %.5f \n", desimal)
	// fmt.Printf("test bool= %t \n", bolecuria)

	var name = "test"
	var e = name[0]
	var eString = string(e)
	fmt.Println(eString)
	fmt.Println(name)
	var values = [...]int{
		1,
		2,
		3,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 25
	fmt.Println(values)
}
