package main

import (
	"errors"
	"fmt"
)

/*
 > Go lang memiliki interface yang digunakan sebagai kontrak untuk
membuat error, yaitu interface error.
> GO lang punya library untuk membuat helper secara mudah,di package errors
*/

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}
func main() {
	// var contohError error = errors.New("UPS ERROR")
	// fmt.Println(contohError.Error())

	hasil, err := Pembagi(5, 0)
	if err == nil {
		fmt.Println("HASILl:", hasil)
	} else {
		fmt.Println("ERROR:", err.Error())
	}
}
