# vogoncal
Simple calendar utility

## Introduction

vogoncal is a simple utility that serves as a reminded of important
dates. You keep the important dates in text files in a simple format:

```
date
text
```

Where date is the date in the format YYYY-MM-DD and text is one or more
lines of description.

For example:

```
2024-12-31
New Year's Eve! Remember to
water the plants and turn off all the lights before going
to the party.
```

## Configuration

Create the configuration file in ~/.vogoncal.cfg.

An example configuration file:

```
rootdir = "/home/user/vogoncal.dir"
calfilepattern = ".*\\.cal"
```

- rootdir - the root directory of the directory tree that contains
            calendar files
- calfilepattern - pattern for calendar files

The pattern in the example is ".*\\.cal", which captures all files with
the extension .cal.

Then keep all calendar files under /home/user/vogoncal.dir in any
subdirectories you wish. For example, make subdirectories 2024, 2025
etc. for each year.

## Calendar file

Every calendar file has zero of more calendar entries, separated with empty line(s), in the format as shown above in the introduction.

## Compiling

go build .

## Usage

```
vogoncal -h

Displays helpful message:

./vogoncal tool. Developed by Aleksandar Janicijevic
Copyright 2024
Usage information:
  -future int
    	How many days into the future
  -m	Print entries for the coming month
  -now string
    	Use as today's date
  -past int
    	How many days into the past
  -w	Print entries for the coming week
  -y	Print entries for the coming year
```

Run vogoncal without any arguments to display entries in the range
from 7 days in the past to 14 days in the future.

Examples:
- Display calendar for the next 20 days:
  vogoncal -future 20
- Display calendar for the days starting on January 1, 2026 and the
  following month:
  vogoncal -now 2026-01-01 -m
- Display calendar for the next year (i.e. today and the next 365 days):
  vogoncal -y

## Wildcards

In the calendar entries, year, month or day can be replaced with "*"
to mean "any". For example:

    2025-*-01

means the first day of each month in 2025.
