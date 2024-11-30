package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello world")
	entries, err := processFile("./test.cal")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Entries: %v\n", entries)
}
