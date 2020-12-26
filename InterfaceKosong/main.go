package main

import "fmt"

/* APA ITU INTERFACE KOSONG ?*/
/*
Di bahasa pemrograman berorientasi objek, ada satu data parent (super class) yang dianggap sebagai
"implementasi" data yang ada di bahasa pemrograman tersebut
cth:
Java=>java.lang.Object
Kotlin=>any
Golang mengatasi tersebut dengan menggunakan "interface kosong"
"Interface Kosong" tidak memiliki deklarasi method satupun,hal ini membuat
secara otomatis semua tipe data akan menjadi implementasinya
*/

/*CONTOH penggunaan Interface Kosong di golang
> fmt.Println(a...interface{})
>panic
>recover
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}

}
func main() {
	//gak bisa bakalan error
	// var data int = Ups(1)
	// fmt.Println(data)

	//type nya harus interface
	var data interface{} = Ups(1)
	fmt.Println(data)
}
