package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	/**
	Pass by value (duplicate)
	*/
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	/**
	Pass by reference with pointer (integrate)
	*/
	address3 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address4 := &address3

	address4.City = "Bandung"

	fmt.Println(address3)
	fmt.Println(address4)

	/**
	Perubahan pada pointer tidak akan merubah variabel reference nya (jika tanpa *)
	*/
	address5 := Address{"Jember", "Jawa Timur", "Indonesia"}
	address6 := &address5

	address6 = &Address{"Denpasar", "Bali", "Indonesia"}
	fmt.Println(address5)
	fmt.Println(address6)

	/**
	Perubahan pada pointer akan merubah variabel reference nya (jika dengan *)
	*/
	address7 := Address{"Salatiga", "Jawa Tengah", "Indonesia"}
	address8 := &address7

	*address8 = Address{"Lombok", "Nusa Tenggara Barat", "Indonesia"}
	fmt.Println(address7)
	fmt.Println(address8)

	/**
	Function new untuk membuat pointer baru yang kosong
	*/
	address9 := new(Address)
	address9.City = "Bandung"
	fmt.Println(address9)

	/**
	Pointer di Function
	*/
	alamat := &Address{"Subang", "Jawa Barat", ""}
	ChangeCountryToIndonesia(alamat)
	fmt.Println(alamat)
}
