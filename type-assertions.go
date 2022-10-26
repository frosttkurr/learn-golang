package main

import "fmt"

func random() interface{} {
	return "Mantap"
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	case bool:
		fmt.Println("Value", value, "is boolean")
	default:
		fmt.Println("Unknown type")
	}
}
