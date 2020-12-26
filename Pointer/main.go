package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	address2 := address1

	address2.City = "Bogor"

	fmt.Println(address1) //tidak berubah karena datanya di copy
	fmt.Println(address2)

	fmt.Println("=======================================")
	/*Implementasi pointer*/
	address3 := &address1
	//line code diatas sama=> var address3 *Address = &address1
	address3.City = "Karawang"

	fmt.Println(address1) // berubah karena datanya di pass by refrence
	fmt.Println(address3)

	/*PENGGUNAAN OPERTAOR '*' */
	fmt.Println("===================OPERATOR * ====================")

	address4 := Address{
		"Surabaya",
		"Jawa Timur",
		"Indonesia",
	}
	address5 := &address4
	address6 := &address4

	//jika ingin mengubah seluruh addressnya
	*address5 = Address{"Jember", "Jawa Timur", "Indonesia"}

	fmt.Println(address4) //adress4 akan dipaksa untuk merefere address 5
	fmt.Println(address5)
	fmt.Println(address6)

	/*PENGGUNAAN FUNCTION NEW */
	fmt.Println("===================FUNCTION NEW ====================")
	//Membuat poi nter baru ke data kosong
	alamat1 := new(Address)
	// alamat1 := new(Address) == var alamat1 *Address	=new (Address)
	alamat1.City = "Merauke"
	fmt.Println(alamat1)
}

/*
>>>>>>>>>>>>>>POINTER<<<<<<<<<<<<<<<<<<<<

#Secara default di Golang semua variable di passing by value bukan by reference.
#POINTER adalah kemampuan untuk membuat reference ke lokasi data di memory
yang sama, tanpa menduplikasi data yang sudah ada.
#kita bisa bikin jadi pass by reference



*/
