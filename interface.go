package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	eko := Person{Name: "Eko"}
	SayHello(eko)

	cat := Animal{Name: "Pus"}
	SayHello(cat)
}
