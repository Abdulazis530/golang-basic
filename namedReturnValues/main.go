package main

import "fmt"

// block kurung kedua adalah informasi megenai variable yang akan direturn, bisa dituliskan satu satu kaya gini ( fistname string, midleName,string, lastName string) tapi karena semua tipe datanya sama tipe data bisa hanya dituliskan sekali diakhir
func getFullName() (firstName, midleName, lastName string) {
	firstName = "abdul"
	midleName = "gokil"
	lastName = "azis"

	//satu hal yang unik dari function named return values adalah kita bisa mendeklarasikan setiap values yang direturn lalu untuk melakukan "return" hanya cukup menuliskan return

	return
}
func main() {
	//cara menggunakannya seperti destructuring di javascript

	// firstName, midleName, lastName := getFullName()
	// fmt.Println(firstName, midleName, lastName)

	//nama variable juga boleh beda dan ga harus sama

	firstName1, midleName2, lastName2 := getFullName()
	fmt.Println(firstName1, midleName2, lastName2)

}
