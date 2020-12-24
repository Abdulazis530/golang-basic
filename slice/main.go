package main

import "fmt"

/*
Tipe Data Slice:
1. memiliki 3 data : pointer, length, dan kapasitas
2. pointer adalah petunjuk data pertama di array pada slice
3. length adalah pajang dari slice
4. kapasitas >= length
*/

//DEMO:
func main() {
	/*Maembuat Array*/
	// month := [...]string{
	// 	"Januari",
	// 	"Februari",
	// 	"Maret",
	// 	"April",
	// 	"Mei",
	// 	"Juni",
	// 	"Juli",
	// 	"Agustus",
	// 	"September",
	// 	"Oktober",
	// 	"November",
	// 	"Desember",
	// }

	/*Maembuat Slice dari Array*/
	// slice := month[10:12]

	// fmt.Println(slice[0])
	// fmt.Println(slice[1])

	/*Funtion pada Slice */

	//1. len mengembalikan panjang dari slice
	// fmt.Println(len(slice))

	//2. cap mengembalikan kapasitas dari slice
	// fmt.Println(cap(slice))

	//3. append => membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru

	// days := [...]string{
	// 	"Senin",
	// 	"Selasa",
	// 	"Rabu",
	// 	"Kamis",
	// 	"Jumat",
	// 	"Sabtu",
	// 	"Minggu"}

	//Skema append jika melebihi kapasitas
	// daySlice1 := days[5:]

	// daySlice1[0] = "Sabtu Lagi"
	// daySlice1[1] = "Minggu Lagi"
	// fmt.Println(days)

	// daySlice2 := append(daySlice1, "Libur lagi")
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	//Skema append jika tidak melebihi kapasitas
	// daySlice1 := days[1:3]
	// fmt.Println(daySlice1)

	// daySlice1[0] = "Sabtu Lagi"
	// daySlice1[1] = "Minggu Lagi"
	// fmt.Println(days)

	// daySlice2 := append(daySlice1, "Libur lagi")
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	//4. make slice => make([]type,length,cap)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "abdul"
	newSlice[1] = "azis"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}
