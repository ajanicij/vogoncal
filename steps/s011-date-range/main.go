package main

import (
	"fmt"
	"log"
)

func check(pattern string, date Date, days int) {
	filter, err := ParseDateFilter(pattern)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DateFilter: %v, Date: %v: ", filter, date)
	pass, dateNew := filter.RangePass(date, days)
	if !pass {
		fmt.Println("FAIL")
	} else {
		fmt.Printf("PASS: first passing date = %v\n", dateNew)
	}
}

func main() {
	pattern := "2024-*-1"
	check(pattern, NewDate(2024, 12, 1), 1)
	check(pattern, NewDate(2024, 12, 2), 1)
	check(pattern, NewDate(2024, 11, 20), 12)
	check(pattern, NewDate(2024, 11, 20), 10)
	check(pattern, NewDate(2024, 11, 21), 11)
	check(pattern, NewDate(2024, 1, 20), 400)

	pattern = "2024-*-29"
	check(pattern, NewDate(2024, 2, 20), 20)
}
