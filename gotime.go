package gotime

import (
	"time"
)

// TODO
// https://pkg.go.dev/time#RFC3339
// 23/25/2014
// 2014-25-23
// 10 06 2014
// 01 Jan 1970 00:00:00 GMT
// 02 Jan 06 15:04 MST
// 2011-10-10T14:48:00  2019-01-01T00:00:00
// 2011-10-10T14:48:00.000+09:00
// January 1, 1970 00:00:00 UTC
// March 7, 2014
// Aug 9, 1995
// Wed, 09 Aug 1995 00:00:00 GMT
// Wed, 09 Aug 1995 00:00:00
// Mon Jan 2 15:04:05 MST 2006

const (
	// 2022/10/12, 2022-10-12, 2020 10 12
	pattern1 = `[0-9]{4}(-|/|[[:space:]])[0-9]{1,2}(-|/|[[:space:]])[0-9]{1,2}`
	// 12/10/2022, 12-10-2022, 25 07 2016
	pattern2 = `[0-9]{1,2}(-|/|[[:space:]])[0-9]{1,2}(-|/|[[:space:]])[0-9]{4}`
	// 01 Jan 1970 00:00:00 GMT
	pattern3 = `[0-9]{1,2} [[:word:]]+ [0-9]{2,4} [0-9]{2}:[0-9]{2}:[0-9]{2} [A-Z]+`
)

// Parse returns time.Time reprsentation of string
func Parse(input string) (parsedTime time.Time, err error) {
	// pattern := "[0-9]{2,4}(-|/)[0-9]{1,2}(-|/)[0-9]{1,2}"
	switch {
	case match(pattern2, input):
		return formatPattern2(input)
	case match(pattern1, input):
		return formatPattern1(input)
	case match(pattern3, input):
		return formatPattern3(input)
	}

	return
}
