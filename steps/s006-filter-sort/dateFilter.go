package main

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
