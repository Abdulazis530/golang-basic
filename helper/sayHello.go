package helper

import "fmt"

//Access Modifier
var version = "2.2"        //huruf kecil tidak bisa diakses dari luar
var Application = "golang" //bisa diakses dari luar

func SayHello(name string) {
	fmt.Println("Hello " + name)
}
