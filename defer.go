package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() //Selalu diawal agar menghindari tidak ke-run saat error
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
