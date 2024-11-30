## s001-parse-date
Parsing date in the form 2024-11-23.

## s002-read-entry
Reading calendar entries from a file.

## s003-date
Type Date and comparing two dates (methods Before and After).

## s004-date-filter
DateFilter that checks if a date is between From and To.

## s005-sort-dates
Sort a slice of dates using sort.Slice.

## s006-filter-sort
Read entries from directory tree, filter by date, sort.

## s007-date-match
Filter is a pattern: 2024, -1, 1 matches a date: year=2024, month=any,
day=1.

## s008-parse-filter
Parse pattern to date filter, e.g. parse "2024-*-01" to
filter: year=2024, month=-1, day=1.

## s009-date-match
Filter is generated from a string pattern, like in s008-parse-filter.
Then a date is matched against the filter.

## s010-date-arithmetic
Adding to or subtracting days from date.

## s011-date-range
Like s009-date-match, but matching filter agains a range of dates.
RangePass gets date d and the number of days n and checks if any of
dates from d to d+n-1 pass the filter.

## s012-comments
Parsing date entries file that can contain comments. Outputting error to
stderr if there is an error in parsing.

## s013-command-line
Parsing command line arguments.

## s014-get-home-dir
Get current user's home directory.

## s015-read-configuration
Read configuration from a file.

## s016-walk-directory
Walk a directory tree.

## s017-read-entries
Walk a directory tree and process calendar files.

TODO:
- Read configuration (entries directory)
- Parse entries directory tree
