package gotime

import (
	"fmt"
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
	timeLocation, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.UTC
	}
	return timeLocation
}

func getTimeZoneInt(hour, minutes int) *time.Location {
	secondsOfUTC := int((time.Duration(hour) * time.Hour).Seconds() + (time.Duration(minutes) * time.Minute).Seconds())
	return time.FixedZone("UTC", secondsOfUTC)
}

// formatPattern1 2022/10/12, 2022-10-12, 2020 10 12
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

// formatPattern2 12/10/2022, 12-10-2022, 25 07 2016
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

// formatPattern3 01 Jan 1970 00:00:00 GMT, 02 Jan 06 15:04 MST
func formatPattern3(stringTime string) (parsedTime time.Time, err error) {
	s := regexp.MustCompile(`[[:space:]]`).Split(stringTime, 5)
	timeZone := getTimeZone(s[len(s)-1])

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

	seconds := 0
	if len(arr) >= 5 {
		seconds = arr[4]
	}

	return time.Date(year, month, arr[0], arr[2], arr[3], seconds, 0, timeZone), nil

}

// formatPattern4 2011-10-10T14:48:00, 2011-10-10T14:48:00 GMT
func formatPattern4(stringTime string) (parsedTime time.Time, err error) {
	s := regexp.MustCompile(`(-|/|[[:space:]])`).Split(stringTime, 4)
	timeZone := getTimeZone(s[len(s)-1])

	list := golist.NewList(s[:2])
	timeArr := strings.Split(s[2], "T")
	list.Append(timeArr[0])

	timeStamp := strings.Split(timeArr[1], ":")
	list, _ = list.Add(golist.NewList(timeStamp))
	arr, err := list.ConvertToSliceInt()
	if err != nil {
		return time.Time{}, err
	}
	year := getYear(arr[0])
	month := time.Month(arr[1])
	day := arr[2]
	hour := arr[3]
	minute := arr[4]

	seconds := 0
	if len(arr) >= 6 {
		seconds = arr[5]
	}

	return time.Date(year, month, day, hour, minute, seconds, 0, timeZone), nil

}

// formatPattern5 2011-10-10T14:48:00.000+09:00
func formatPattern5(stringTime string) (parsedTime time.Time, err error) {
	s := regexp.MustCompile(`(T|-|/|[[:space:]]|:|\.|\+)`).Split(stringTime, 10)

	list := golist.NewList(s)

	arr, err := list.ConvertToSliceInt()
	if err != nil {
		return time.Time{}, err
	}
	year := getYear(arr[0])
	month := time.Month(arr[1])
	day := arr[2]
	hour := arr[3]
	minute := arr[4]
	seconds := arr[5]
	nseconds := arr[6]
	timeZone := getTimeZoneInt(arr[7], arr[8])

	return time.Date(year, month, day, hour, minute, seconds, nseconds, timeZone), nil

}

// formatPattern6 January 1, 1970 00:00:00 UTC, January 1, 09 00:00:00.00 GMT
func formatPattern6(stringTime string) (parsedTime time.Time, err error) {
	s := regexp.MustCompile(`(\.|[[:space:]]|:|(,[[:space:]]))`).Split(stringTime, 9)
	month := getMonth(s[0])
	timeZone := getTimeZone(s[len(s)-1])
	list := golist.NewList(s[1 : len(s)-1])
	fmt.Println(golist.NewList(s))

	arr, err := list.ConvertToSliceInt()
	if err != nil {
		return time.Time{}, err
	}
	year := getYear(arr[1])

	day := arr[0]
	hour := arr[2]
	minute := arr[3]
	seconds := 0
	nseconds := 0
	if len(arr) >= 5 {
		seconds = arr[4]
	}
	if len(arr) >= 6 {
		nseconds = arr[5]
	}

	return time.Date(year, month, day, hour, minute, seconds, nseconds, timeZone), nil

}
