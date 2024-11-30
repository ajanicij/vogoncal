package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	now := flag.String("now", "", "Use as today's date")
	future := flag.Int("future", 0, "How many days into the future")
	past := flag.Int("past", 0, "How many days into the past")

	week := flag.Bool("w", false, "Print entries for the coming week")
	month := flag.Bool("m", false, "Print entries for the coming month")
	year := flag.Bool("y", false, "Print entries for the coming year")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"\n%s tool. Developed by Aleksandar Janicijevic\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2024\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Count flags in order to determine if the command is used correctly.
	// - We don't allow future or past if week, month or year is true.
	// - Only one of week, month or year can be true.
	countFlags := 0
	if *week {
		countFlags++
	}
	if *month {
		countFlags++
	}
	if *year {
		countFlags++
	}
	if countFlags > 1 {
		flag.Usage()
		os.Exit(1)
	}
	if (*future > 0 || *past > 0) && countFlags == 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *week {
		*future = 6
	}
	if *month {
		*future = 31 // One more, just in case
	}
	if *year {
		*future = 366 // One more, in case of a leap year
	}

	var t time.Time
	var date Date
	var err error
	if *now == "" {
		t = time.Now()
		date = getDate(t)
	} else {
		date, err = ParseDate(*now)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("now: %s\n", *now)
	fmt.Printf("future: %d\n", *future)
	fmt.Printf("past: %d\n", *past)
	fmt.Printf("date: %v\n", date)

	startDate, days := date.Range(*future, *past)
	fmt.Printf("range: %v %d\n", startDate, days)
}
