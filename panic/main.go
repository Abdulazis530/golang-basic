package main

import "fmt"

func endApp() {
	fmt.Println("Applikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERRORR")
	}
	fmt.Println("Applikasi berjalan")
}
func main() {
	runApp(false)
	runApp(true)
}
