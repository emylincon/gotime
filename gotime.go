package gotime

import (
	"time"
)

// TODO
// https://pkg.go.dev/time#RFC3339

const (
	// 2022/10/12, 2022-10-12, 2020 10 12
	pattern1 = `^[0-9]{4}(-|/|[[:space:]])[0-9]{1,2}(-|/|[[:space:]])[0-9]{1,2}$`
	// 12/10/2022, 12-10-2022, 25 07 2016
	pattern2 = `^[0-9]{1,2}(-|/|[[:space:]])[0-9]{1,2}(-|/|[[:space:]])[0-9]{4}$`
	// 01 Jan 1970 00:00:00 GMT, 02 Jan 06 15:04 MST
	pattern3 = `^[0-9]{1,2} [[:word:]]+ [0-9]{2,4} [0-9]{2}:[0-9]{2}(:[0-9]{2}){0,1} ([A-Z]+){0,1}$`
	// 2011-10-10T14:48:00, 2011-10-10T14:48:00 GMT
	pattern4 = `^[0-9]{2,4}(-|/|[[:space:]])[0-9]{1,2}(-|/|[[:space:]])[0-9]{1,2}T[0-9]{2}:[0-9]{2}(:[0-9]{2}){0,1}([[:space:]][A-Z]+){0,1}$`
	// 2011-10-10T14:48:00.000+09:00
	pattern5 = `^[0-9]{2,4}(-|/|[[:space:]])[0-9]{1,2}(-|/|[[:space:]])[0-9]{1,2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]+\+[0-9]{2}:[0:9]{2}$`
	// January 1, 1970 00:00:00 UTC, January 1, 09 00:00:00.00 GMT, August 7, 2014 or Aug 7, 2014
	pattern6 = `[[:word:]]+ [0-9]{1,2}, [0-9]{2,4}([[:space]][0-9]{2}:[0-9]{2}(:[0-9]{2}){0,1}(\.[0-9]+){0,1}([[:space]][A-Z]+){0,1}){0,1}`
	// Wed, 09 Aug 1995 00:00:00 GMT or Wed, 09 Aug 1995 00:00:00
	pattern7 = `[[:word:]]+, [0-9]{1,2} [[:word:]]+ [0-9]{2,4}([[:space]][0-9]{2}:[0-9]{2}(:[0-9]{2}){0,1}(\.[0-9]+){0,1}([[:space]][A-Z]+){0,1}){0,1}`
	// Mon Jan 2 15:04:05 MST 2006
	pattern8 = `^[[:word:]]+ [[:word:]]+ [0-9]{1,2} [0-9]{2}:[0-9]{2}:[0-9]{2} [A-Z]+ [0-9]{2,4}$`
)

// Parse returns time.Time reprsentation of string
func Parse(input string) (parsedTime time.Time, err error) {
	switch {
	case match(pattern5, input):
		return formatPattern5(input)
	case match(pattern4, input):
		return formatPattern4(input)
	case match(pattern2, input):
		return formatPattern2(input)
	case match(pattern1, input):
		return formatPattern1(input)
	case match(pattern3, input):
		return formatPattern3(input)
	case match(pattern8, input):
		return formatPattern8(input)
	case match(pattern6, input):
		return formatPattern6(input)
	case match(pattern7, input):
		return formatPattern7(input)

	}
	return
}
