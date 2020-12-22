package main

import "fmt"

func main() {
	var namaDepan string = "Azis"
	var namaBelakang string
	namaBelakang = "gokil"

	//variable tanpa typedata
	namangasal := "aduh"
	var namangasal2 = "test"

	//variable tempat sampah
	// _ = "sampah"

	//di go bisa kaya gini
	namaHantu, _ := "pocong", "sampah"

	fmt.Printf("halo %s %s \n", namaDepan, namaBelakang)
	fmt.Println("halo" + namaBelakang + " " + namaDepan + " " + namangasal + " " + namangasal2)
	fmt.Println("halo" + namaHantu)
	// fmt.Println("halo" + _) variable sampah ga bisa ditampilin
}
