package main

import (
	"fmt"
	"log"
)

func checkFilter(pattern string) {
	filter, err := ParseDateFilter(pattern)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pattern: %s, filter: %v\n", pattern, filter)
}

func main() {
	checkFilter("2024-12-1")
	checkFilter("2024-12-*")
	checkFilter("2024-12-01")
	checkFilter("2024-*-1")
	checkFilter("*-12-1")
	checkFilter("*-*-1")
}
