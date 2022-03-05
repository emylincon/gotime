package gotime

import (
	"regexp"
	"strings"
	"time"

	"github.com/emylincon/golist"
)

func match(pattern, input string) bool {
	boolean, _ := regexp.MatchString(pattern, input)
	return boolean
}

func getYear(year int) int {
	if year >= 69 && year <= 99 { // Unix time starts Dec 31 1969 in some time zones
		year += 1900
	} else if year < 69 {
		year += 2000
	}
	return year

}

func getMonth(month string) time.Month {
	switch strings.ToLower(month)[:3] {
	case "jan":
		return time.January
	case "feb":
		return time.February
	case "mar":
		return time.March
	case "apr":
		return time.April
	case "may":
		return time.May
	case "jun":
		return time.June
	case "jul":
		return time.July
	case "aug":
		return time.August
	case "sep":
		return time.September
	case "oct":
		return time.October
	case "nov":
		return time.November
	case "dec":
		return time.December
	}
	return time.January
}

func getTimeZone(timeZone string) *time.Location {
	if timeZone == "GMT" {
		return time.UTC
	}
	return time.Local
}

func formatPattern1(stringTime string) (parsedTime time.Time, err error) {
	s := regexp.MustCompile(`(-|/|[[:space:]])`).Split(stringTime, 3)
	list := golist.NewList(s)
	arr, err := list.ConvertToSliceInt()
	if err != nil {
		return time.Time{}, err
	}
	year := getYear(arr[0])

	return time.Date(year, time.Month(arr[1]), arr[2], 0, 0, 0, 0, time.UTC), nil

}

func formatPattern2(stringTime string) (parsedTime time.Time, err error) {
	s := regexp.MustCompile(`(-|/|[[:space:]])`).Split(stringTime, 3)
	list := golist.NewList(s)
	arr, err := list.ConvertToSliceInt()
	if err != nil {
		return time.Time{}, err
	}
	year := getYear(arr[2])

	return time.Date(year, time.Month(arr[1]), arr[0], 0, 0, 0, 0, time.UTC), nil

}

func formatPattern3(stringTime string) (parsedTime time.Time, err error) {
	s := regexp.MustCompile(`[[:space:]]`).Split(stringTime, 5)
	list := golist.NewList([]string{s[0]})
	list.Append(s[2])
	timeStamp := strings.Split(s[3], ":")
	list, _ = list.Add(golist.NewList(timeStamp))
	arr, err := list.ConvertToSliceInt()
	if err != nil {
		return time.Time{}, err
	}
	year := getYear(arr[1])
	month := getMonth(s[1])
	timeZone := getTimeZone(s[len(s)-1])

	return time.Date(year, month, arr[0], arr[2], arr[3], arr[4], 0, timeZone), nil

}
