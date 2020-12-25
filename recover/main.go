package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)

	}
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
	// runApp(true)
}

//Recover adalah function yag bisa kita gunakan untuk menangkap data panic dengan recover proses panic akan terhenti dan program akan terus jalan

//Hal penting:
//fungsi dibawaha panic tidak akan pernah di eksekusi oleh arena itu Recover diletekan didalam "fungsi yang di Defer"
