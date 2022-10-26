package main

/**
Add _ in front of package if not used but don't want to be erased
Eg. _ import "learn-golang/helper"
*/
import (
	"fmt"
	"learn-golang/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
