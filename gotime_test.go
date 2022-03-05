package gotime_test

import (
	"testing"
	"time"

	"github.com/emylincon/gotime"
)

func getTimeZone(timeZone string) *time.Location {
	timeLocation, err := time.LoadLocation(timeZone)
	if err != nil {
		return time.Local
	}
	return timeLocation
}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected time.Time
	}{
		{
			desc:     "2012 07 12, pattern1",
			input:    "2012 07 12",
			expected: time.Date(2012, time.Month(7), 12, 0, 0, 0, 0, time.UTC),
		},
		{
			desc:     "12 07 2012, pattern2",
			input:    "12 07 2012",
			expected: time.Date(2012, time.Month(7), 12, 0, 0, 0, 0, time.UTC),
		},
		{
			desc:     "01 January 1970 00:00:00 GMT, pattern3 1",
			input:    "01 January 1970 00:00:00 GMT",
			expected: time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, getTimeZone("GMT")),
		},
		{
			desc:     "01 January 70 00:00 GMT, pattern3 2",
			input:    "01 January 70 00:00 GMT",
			expected: time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, getTimeZone("GMT")),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := gotime.Parse(tC.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got.String() != tC.expected.String() {
				t.Errorf("Error: got : %v != expected : %v", got, tC.expected)
			}
		})
	}
}
