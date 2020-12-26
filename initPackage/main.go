package main

// import (
// 	"basic/initPackage/database"
// 	"fmt"
// )

//Black identifier digunakan jika kita mengimpor package namun package tersebut tidak diakases

import (
	_ "basic/initPackage/database"
	"fmt"
)

func main() {
	// connection := database.GetDataBase()
	fmt.Println("blank identifier")
}
