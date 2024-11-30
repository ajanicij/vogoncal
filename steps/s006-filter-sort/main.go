package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// Calendar entry
type Entry struct {
	Date Date
	Text []string
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dateFilter := NewDateFilter(
		NewDate(2024, 11, 23),
		NewDate(2024, 12, 15),
	)

	allEntries := []Entry{}

	dir := "./testdir"
	// Walk through the directory tree from ./testdir.
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var entries []Entry
		if !info.IsDir() {
			entries, err = readEntries(path, dateFilter)
		}
		if err != nil {
			return err
		}

		allEntries = append(entries, allEntries...)

		return err
	})
	if err != nil {
		return err
	}
	fmt.Printf("filtered entries: %v\n", allEntries)
	SortEntries(allEntries)
	fmt.Printf("Sorted entries: %v\n", allEntries)
	return nil
}

func readEntries(filename string, filter DateFilter) ([]Entry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	entries := []Entry{}

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

		dateTime, err := time.Parse("2006-01-02", datestr)
		if err != nil {
			continue
		}

		date := NewDate(dateTime.Year(), int(dateTime.Month()), dateTime.Day())

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

		entry := Entry{
			Date: date,
			Text: lines,
		}

		fmt.Printf("Entry: %v\n", entry)

		if filter.Pass(date) {
			fmt.Println("Appending entry")
			entries = append(entries, entry)
		} else {
			fmt.Println("Not appending entry")
		}

		if eof {
			break
		}
	}

	return entries, nil

}

func SortEntries(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool {
		date1 := entries[i].Date
		date2 := entries[j].Date
		return date1.Before(date2)
	})
}
