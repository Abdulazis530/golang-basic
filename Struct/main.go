package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var azis Customer
	azis.Name = "Abdul Azis"
	azis.Address = "Turkey"
	azis.Age = 25

	fmt.Println(azis)
	fmt.Println(azis.Name)

	//Struct Literal
	//Struct Literal 1 :
	firaun := Customer{
		Name:    "Firaun",
		Address: "Mesir",
		Age:     22222,
	}

	fmt.Println(firaun)
	//Struct Literal 2:
	//not recomended rawan eror
	roroKidul := Customer{
		"roroKidul",
		"pantai selatan",
		22223,
	}

	fmt.Println(roroKidul)

}

/*
Struct mirip class
Struct => template data
*/
