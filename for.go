package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}

	slice := []string{
		"Eko",
		"Kurniawan",
		"Khannedy",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, value := range slice {
		fmt.Println("Index", index, "=", value)
	}

	for _, value := range slice { // _ digunakan agar variable tidak dianggap
		// fmt.Println("Index", index, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}