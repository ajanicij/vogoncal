package main

import (
	"fmt"
	"log"
)

func check(pattern string, date Date) {
	filter, err := ParseDateFilter(pattern)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DateFilter: %v, Date: %v: ", filter, date)
	if filter.Pass(date) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

func main() {
	pattern := "2024-*-1"
	check(pattern, NewDate(2024, 12, 1))
	check(pattern, NewDate(2024, 3, 1))
	check(pattern, NewDate(2024, 12, 13))
	check(pattern, NewDate(2025, 12, 1))
}
