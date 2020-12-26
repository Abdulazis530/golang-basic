package main

import "fmt"

/*

#type Assertsions adalah kemampuan merubah tipe data menjadai
tipe data yang diinginkan

#Fitur ini sering digunakan ketika bertemu dengan tipe interface
kosong

# agar lebih aman menggunakan switch jika ingin type assetsions guna
mencegah aplikasi berrhenti akibat panic
*/

func random() interface{} {
	return 5
}

func main() {
	// var result interface{} = random()
	// var resultString string = result.(string)

	// fmt.Println(resultString)

	//cth salah

	// var resultNumber int = result.(int) //panic
	// fmt.Println(resultNumber)

	//MENGGUNAKAN SWITCH
	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "is string")
	case int:
		fmt.Println("value", value, "is integer")

	default:
		fmt.Println("Unknown type")
	}

}
