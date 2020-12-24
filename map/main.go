package main

import "fmt"

/*
Map adalah tipe data kumpulan key-value dengan tipe data yang sama (mirip object di javascript tapi kaya array)
*/
func main() {
	/*
		hal penting pada Map
		1.len(map)
		2.map[key] mengambil data pada map
		3.map[key]=value mengubah data di map dengan key
		4.make(map[typekey]typevalue) => syntax untuk membuat map baru
		5.delete(map,key) => menghapus data di map dengan key
	*/

	//deklarasi panjang
	/* var person map[string]string = map[string]string{
		"name":    "Joji",
		"address": "Los Santos",
	}
	*/

	//deklarasi pendek
	person := map[string]string{
		"name":    "Joji",
		"address": "Los Santos",
	}

	//tambah data
	person["job"] = "Kuli"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["job"])
	delete(person, "job")
	fmt.Println(person)

}
