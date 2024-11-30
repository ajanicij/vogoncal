package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

// Type Date for operations on dates, disregarding
// the time within the day
type Date struct {
	Year  int
	Month int
	Day   int
}

// Calendar entry
type Entry struct {
	Date Date
	Text []string
}

func NewDate(year, month, day int) Date {
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func getDate(t time.Time) Date {
	return Date{
		Year:  t.Year(),
		Month: int(t.Month()),
		Day:   t.Day(),
	}
}

func (d Date) Before(other Date) bool {
	t1 := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(other.Year, time.Month(other.Month), other.Day, 0, 0, 0, 0, time.UTC)
	return t1.Before(t2)
}

func (d Date) After(other Date) bool {
	t1 := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(other.Year, time.Month(other.Month), other.Day, 0, 0, 0, 0, time.UTC)
	return t1.After(t2)
}

func (d Date) AddDays(n int) Date {
	t1 := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	t2 := t1.AddDate(0, 0, n)
	return Date{
		Year:  t2.Year(),
		Month: int(t2.Month()),
		Day:   t2.Day(),
	}
}

func SortDates(dates []Date) {
	sort.Slice(dates, func(i, j int) bool {
		date1 := dates[i]
		date2 := dates[j]
		return date1.Before(date2)
	})
}

func processFile(path string) ([]Entry, error) {
	entries := []Entry{}
	file, err := os.Open(path)
	if err != nil {
		return entries, err
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
		if datestr == "" || datestr[0] == '#' {
			continue
		}

		dateTime, err := time.Parse("2006-01-02", datestr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %s as date: %s\n",
				datestr, err.Error())
			continue
		}

		date := Date{
			Year:  dateTime.Year(),
			Month: int(dateTime.Month()),
			Day:   dateTime.Day(),
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

		entry := Entry{
			Date: date,
			Text: lines,
		}

		fmt.Printf("Entry: %v\n", entry)
		entries = append(entries, entry)

		if eof {
			break
		}
	}

	return entries, nil
}
