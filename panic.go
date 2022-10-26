package main

import "fmt"

func endApp() {
	message := recover() //Recover dari error
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp() //Defer to run code in the end - position always first
	if error {
		panic("APLIKASI ERROR") //Panic for error
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
