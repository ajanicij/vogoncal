package main

import (
	"fmt"
	"log"
	"os"

	"path/filepath"
)

func main() {
	fmt.Println("hello world")

	err := filepath.Walk("./dir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("Walk: entry: %s\n", path)
		if info.IsDir() {
			fmt.Println("  -- directory")
		} else {
			fmt.Println("  -- file")
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
