package main

import (
	"fmt"
	"time"
)

// Type Data for operations on dates, disregarding
// the time within the day
type Date struct {
	Year  int
	Month int
	Day   int
}

// Date

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

// DateFilter

type DateFilter struct {
	From Date
	To   Date
}

func NewDateFilter(from, to Date) DateFilter {
	return DateFilter{
		From: from,
		To:   to,
	}
}

func (df DateFilter) Pass(d Date) bool {
	if df.From.After(d) {
		return false
	}
	if df.To.Before(d) {
		return false
	}
	return true
}

func main() {
	dateFilter := NewDateFilter(
		NewDate(2024, 11, 23),
		NewDate(2024, 12, 1),
	)

	date1 := NewDate(2024, 11, 23)
	if !dateFilter.Pass(date1) {
		fmt.Printf("NOT PASS: %v\n", date1)
	}

	date2 := NewDate(2024, 11, 24)
	if !dateFilter.Pass(date2) {
		fmt.Printf("NOT PASS: %v\n", date2)
	}

	date3 := NewDate(2024, 12, 1)
	if !dateFilter.Pass(date3) {
		fmt.Printf("NOT PASS: %v\n", date3)
	}

	// NOT PASS: after the filter
	date4 := NewDate(2024, 12, 2)
	if !dateFilter.Pass(date4) {
		fmt.Printf("NOT PASS: %v\n", date4)
	}

	// NOT PASS: before the filter
	date5 := NewDate(2024, 11, 22)
	if !dateFilter.Pass(date5) {
		fmt.Printf("NOT PASS: %v\n", date5)
	}
}
