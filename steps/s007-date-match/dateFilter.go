package main

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
