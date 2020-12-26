package main

import "fmt"

type Man struct {
	Name string
}

func (m *Man) Married() {
	m.Name = "Mr." + m.Name
}

func main() {
	azis := Man{"Azis"}
	azis.Married()
	fmt.Println(azis)

}
