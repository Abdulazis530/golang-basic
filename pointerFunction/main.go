package main

import "fmt"

/*Saat kita mengubah parameter di function, secara default
adalah pass by value. Jadi data akan di copy lalu dikirim ke function terebut.
jika kita mengubah data didalam function, data yang asli tidak pernah berubah
jika mau mengubah operator bintang bisa digunakan didalam parameter
*/

type Address struct {
	City, Province, Country string
}

func ChangeCountryToUSA(a *Address) {
	a.Country = "USA"
}

func main() {
	address1 := Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	// ChangeCountryToUSA(address1)

	fmt.Println(address1) //alamat tidak berubah

	ChangeCountryToUSA(&address1)
	fmt.Println(address1) //alamat tidak berubah

}
