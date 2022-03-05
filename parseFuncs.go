package gotime

import (
	"regexp"
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
