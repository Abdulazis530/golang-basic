package main

import "fmt"

/*A. ekilas tentang Interface*/
/*Interface adalah tipe data abstrak, tidak memiliki
implementasi langsung, hanya berisi definisi-definisi method
biasanya digunakan sebagai kontrak
*/

/*B.Implementasi Interface*/

/*
Setiap tipe data yang sesuai dengan kontrak interface, secara otamatis
dianggap sebagai Interface. Bahasa lain tidak otomatis jadi harus manual
*/
//Membuat kontrak HasName

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	var azis Person
	azis.Name = "Azis"
	sayHello(azis)

	cat := Animal{
		Name: "Joni",
	}
	sayHello(cat)

}
