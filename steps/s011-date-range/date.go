package main

import (
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
