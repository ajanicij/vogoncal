package main

import (
	"fmt"
	"strconv"
	"strings"
)

// DateFilter

type DateFilter struct {
	DatePattern Date
}

func NewDateFilter(year, month, day int) DateFilter {
	return DateFilter{
		DatePattern: Date{
			Year:  year,
			Month: month,
			Day:   day,
		},
	}
}

func ParseDateFilter(pattern string) (DateFilter, error) {
	pattern = strings.TrimRight(pattern, " ")
	list := strings.Split(pattern, "-")
	var filter DateFilter
	if len(list) != 3 {
		return filter, fmt.Errorf("bad pattern: %s", pattern)
	}

	// Parse year.
	patternYear := list[0]
	if patternYear == "*" {
		filter.DatePattern.Year = -1
	} else {
		yearVal, err := strconv.ParseInt(patternYear, 10, 64)
		if err != nil {
			return filter, fmt.Errorf("Error parsing year `%s': %s", patternYear, err.Error)
		}
		filter.DatePattern.Year = int(yearVal)
	}

	// Parse month.
	patternMonth := list[1]
	if patternMonth == "*" {
		filter.DatePattern.Month = -1
	} else {
		monthVal, err := strconv.ParseInt(patternMonth, 10, 64)
		if err != nil {
			return filter, fmt.Errorf("Error parsing month `%s': %s", patternMonth, err.Error)
		}
		filter.DatePattern.Month = int(monthVal)
	}

	// Parse day.
	patternDay := list[2]
	if patternDay == "*" {
		filter.DatePattern.Day = -1
	} else {
		dayVal, err := strconv.ParseInt(patternDay, 10, 64)
		if err != nil {
			return filter, fmt.Errorf("Error parsing day `%s': %s", patternDay, err)
		}
		filter.DatePattern.Day = int(dayVal)
	}

	return filter, nil
}

func (df DateFilter) Pass(d Date) bool {
	if df.DatePattern.Year != -1 && df.DatePattern.Year != d.Year {
		return false
	}
	if df.DatePattern.Month != -1 && df.DatePattern.Month != d.Month {
		return false
	}
	if df.DatePattern.Day != -1 && df.DatePattern.Day != d.Day {
		return false
	}
	return true
}

func (df DateFilter) RangePass(d Date, days int) []Date {
	var dates []Date

	if days <= 0 {
		return dates
	}

	date1 := d
	for i := 0; i < days; i++ {
		if df.Pass(date1) {
			dates = append(dates, date1)
		}
		date1 = date1.AddDays(1)
	}
	return dates
}
