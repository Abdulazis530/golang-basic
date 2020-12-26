package main

import "fmt"

/*
Cara membuat data kosong
namun hanya bisa digunakan di beberapa tipe data
1. Interface
2. Function
3. Map
4. Slice
5. Pointer
6. Channel
*/

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	//cth: map dengan data kosong
	// var person map[string]string = nil
	// fmt.Println(person)

	// person := newMap("Azis")
	// fmt.Println(person)

	// nilPerson := newMap("")
	// fmt.Println(nilPerson)

	//cth: pengecekan menggunakan nil
	var person map[string]string = newMap("")

	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}

}
