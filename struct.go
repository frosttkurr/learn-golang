package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko Kurniawan"
	eko.Address = "Indonesia"
	eko.Age = 30
	eko.Married = true

	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)
	fmt.Println(eko.Married)

	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30, true}
	fmt.Println(budi)

	eko.sayHi("Joko")
}
