package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	date := NewDate(2024, 12, 1)
	fmt.Printf("date: %v\n", date)

	date2 := date.AddDays(7)
	fmt.Printf("date + 7: %v\n", date2)

	date3 := date.AddDays(31)
	fmt.Printf("date + 31: %v\n", date3)

	date4 := date.AddDays(-1)
	fmt.Printf("date - 1: %v\n", date4)
}
