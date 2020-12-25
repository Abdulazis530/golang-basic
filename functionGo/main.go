package main

import "fmt"

//Jika function dijadikan parameter untuk function lain
//maka function tersebut juga harus didefinisikan di parameter "function yang menggunakan" function tersebut

// func censor(name string, filter func(string) string) {
// 	fmt.Println("Hello", filter(name))
// }

//type declaration juga bisa digunakan untuk mendifiniskan type dari filter yang terlalu panjang jadi lebih mudah gitu loh

type Filter func(string) string

func censor(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}
func unapropiateWordFilter(name string) string {
	if name == "Babi" {
		return "****"
	} else {
		return name
	}
}
func main() {
	censor("azis", unapropiateWordFilter)
	//output :Hello azis

	//function juga bisa disimpan didalam variable
	filter := unapropiateWordFilter
	censor("Babi", filter)
	//output :Hello ****

}
