package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("Hei ga boleh ngomong", name)
	} else {
		fmt.Println("welcome", name)
	}

}
func main() {
	blacklist := func(name string) bool {
		return name == "kasar"
	}

	registerUser("azis", blacklist)
	//output : welcome azis
	registerUser("kasar", func(name string) bool {
		return name == "kasar"
	})
	//output : Hei ga boleh ngomong kasar

}
