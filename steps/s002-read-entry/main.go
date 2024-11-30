package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// Calendar entry
type Entry struct {
	Date time.Time
	Text []string
}

func main() {
	fmt.Println("hello world")
	err := processFile("./test.cal")
	if err != nil {
		log.Fatal(err)
	}
}

func processFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for true {
		// Scan date.
		if !scanner.Scan() {
			break
		}

		datestr := scanner.Text()
		// Skip empty lines.
		if datestr == "" {
			continue
		}

		date, err := time.Parse("2006-01-02", datestr)
		if err != nil {
			continue
		}

		eof := false
		lines := []string{}
		for true {
			if !scanner.Scan() {
				eof = true
				break
			}
			line := scanner.Text()
			if line == "" {
				break
			}
			lines = append(lines, line)
		}

		entry := &Entry{
			Date: date,
			Text: lines,
		}

		fmt.Printf("Entry: %v\n", *entry)

		if eof {
			break
		}
	}

	return nil
}
