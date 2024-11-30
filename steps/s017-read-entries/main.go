package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"path/filepath"
)

func GetEntries(path string) ([]Entry, error) {
	allEntries := []Entry{}

	err := filepath.Walk("./testdir", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		matched, err := regexp.Match(".*\\.cal", []byte(path))
		if err != nil {
			return nil
		}

		if !matched {
			fmt.Printf("Skipping file %s\n", path)
			return nil
		}

		fmt.Printf("Processing file %s\n", path)

		entries, err := processFile(path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("entries: %v\n", entries)

		allEntries = append(allEntries, entries...)

		return nil
	})
	if err != nil {
		return allEntries, err
	}
	return allEntries, nil
}

func main() {
	fmt.Println("hello world")

	allEntries, err := GetEntries("./testdir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAll entries: %v\n", allEntries)
}
