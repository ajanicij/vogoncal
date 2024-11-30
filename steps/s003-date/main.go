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

func main() {
	date := NewDate(2024, 11, 23)
	fmt.Printf("Date: %v\n", date)

	date1 := NewDate(2024, 11, 23)

	date2 := NewDate(2025, 1, 1)

	if date1.Before(date2) {
		fmt.Println("Before")
	}
}
