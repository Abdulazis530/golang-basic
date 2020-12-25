package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAplication() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runAplication()
}

//Defer adalah sebuah function yang bisa dijadwalkan pengeksekusiannnya setelah sebuah function selesai di eksekusi
//Selalu dieksekusi walaupun terjadi error di fuction yang dieksekusi
