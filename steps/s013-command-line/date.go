package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

// Range returns start date of the range and the number of days
// in the range.
func (d Date) Range(future, past int) (Date, int) {
	if future < 0 || past < 0 {
		return d, 0
	}
	startDate := d.AddDays(-past)
	return startDate, future + past + 1
}

func SortDates(dates []Date) {
	sort.Slice(dates, func(i, j int) bool {
		date1 := dates[i]
		date2 := dates[j]
		return date1.Before(date2)
	})
}

func ParseDate(datestr string) (Date, error) {
	var date Date
	list := strings.Split(datestr, "-")
	if len(list) != 3 {
		return date, fmt.Errorf("Bad date: %s", datestr)
	}

	// Parse year.
	year, err := strconv.ParseInt(list[0], 10, 64)
	if err != nil {
		return date, fmt.Errorf("Error parsing year %s: %s",
			list[0], err.Error())
	}

	// Parse month.
	month, err := strconv.ParseInt(list[1], 10, 64)
	if err != nil {
		return date, fmt.Errorf("Error parsing month %s: %s",
			list[1], err.Error())
	}

	// Parse month.
	day, err := strconv.ParseInt(list[2], 10, 64)
	if err != nil {
		return date, fmt.Errorf("Error parsing day %s: %s",
			list[2], err.Error())
	}

	return Date{
		Year:  int(year),
		Month: int(month),
		Day:   int(day),
	}, nil
}
