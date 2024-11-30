package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type config struct {
	RootDir        string
	CalFilePattern string
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
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
		fmt.Fprintf(flag.CommandLine.Output(), "Source code: [vogoncal](https://github.com/ajanicij/vogoncal)")
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

	// If none of -future, -past, -w, -m, -y are on the command
	// line, then set the future to 14 days and past to 7 days.
	if *future == 0 && *past == 0 && countFlags == 0 {
		*future = 14
		*past = 7
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

	startDate, days := date.Range(*future, *past)

	// Get home directory.
	dirname, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// fmt.Printf("home directory: %s\n", dirname)

	// Get path to the configuration file.
	file := ".vogoncal.cfg"
	cfgpath := filepath.Join(dirname, file)

	// fmt.Printf("cfgpath is %s\n", cfgpath)

	// Read configuration.
	buf, err := os.ReadFile(cfgpath)
	if err != nil {
		return err
	}
	var cfg config
	err = toml.Unmarshal(buf, &cfg)
	if err != nil {
		return err
	}

	// fmt.Printf("rootdir is %s\n", cfg.RootDir)
	// fmt.Printf("calfilepattern is %s\n", cfg.CalFilePattern)

	allEntries, err := GetEntries(cfg.RootDir, cfg.CalFilePattern, startDate, days)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("all entries: %v\n", allEntries)
	for _, entry := range allEntries {
		// fmt.Printf("entry: %v\n", entry)
		PrintEntry(entry)
	}

	// TODO: sort entries.

	return nil
}

func GetEntries(path string, pattern string, d Date, days int) ([]Entry, error) {
	allEntries := []Entry{}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		matched, err := regexp.Match(pattern, []byte(path))
		if err != nil {
			return nil
		}

		if !matched {
			// fmt.Printf("Skipping file %s\n", path)
			return nil
		}

		// fmt.Printf("Processing file %s\n", path)

		entries, err := ProcessFile(path, d, days)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("entries: %v\n", entries)

		allEntries = append(allEntries, entries...)

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Sort entries by date.
	SortEntries(allEntries)
	return allEntries, err
}

func SortEntries(entries []Entry) {
	sort.Slice(entries, func(i, j int) bool {
		date1 := entries[i].Date
		date2 := entries[j].Date
		return date1.Before(date2)
	})
}

func PrintEntry(entry Entry) {
	// Print date.
	date := time.Date(entry.Date.Year, time.Month(entry.Date.Month), entry.Date.Day, 0, 0, 0, 0, time.UTC)
	datestr := date.Format("2006 January 02")
	fmt.Println("----")
	fmt.Printf("- %s\n", datestr)

	// Print text.
	for _, line := range entry.Text {
		fmt.Printf("%s\n", line)
	}
}
