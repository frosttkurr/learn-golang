package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("eko", blackList)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("eko", func(name string) bool {
		return name == "eko"
	})
}