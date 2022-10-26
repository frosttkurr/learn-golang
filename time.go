package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Local())
	// fmt.Println(now.Date())
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Day())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 10, 27, 02, 10, 02, 12, time.UTC)
	fmt.Println(utc.Local())

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-10-27")
	fmt.Println(parse)
}
