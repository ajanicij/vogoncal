package main

import (
	"fmt"
)

func check(filter DateFilter, date Date) {
	fmt.Printf("DateFilter: %v, Date: %v: ", filter, date)
	if filter.Pass(date) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

func main() {
	fmt.Println("hello world")
	filter := NewDateFilter(2024, -1, 1)
	check(filter, NewDate(2024, 12, 1))
	check(filter, NewDate(2024, 3, 1))
	check(filter, NewDate(2024, 12, 13))
	check(filter, NewDate(2025, 12, 1))
}
