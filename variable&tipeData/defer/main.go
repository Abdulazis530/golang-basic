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
