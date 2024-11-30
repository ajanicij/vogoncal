package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")

	// Get home directory.
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("home directory: %s\n", dirname)
}
