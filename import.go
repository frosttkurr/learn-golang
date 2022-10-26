package main

import (
	"fmt"
	"learn-golang/helper"
)

func main() {
	helper.SayHello("Eko")
	// helper.sayGoodbye("Eko") //can't access from other package if camel case
	fmt.Println(helper.Application) //can access from other package if upper case
}
