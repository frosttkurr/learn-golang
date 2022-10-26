package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian tidak boleh NOL")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	contohError := errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
