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

func getTimeZoneInt(hour, minutes int) *time.Location {
	secondsOfUTC := int((time.Duration(hour) * time.Hour).Seconds() + (time.Duration(minutes) * time.Minute).Seconds())
	return time.FixedZone("UTC", secondsOfUTC)
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
		{
			desc:     "02 Jan 06 15:04 MST, pattern3 3",
			input:    "02 Jan 06 15:04 MST",
			expected: time.Date(2006, time.Month(1), 2, 15, 4, 0, 0, getTimeZone("MST")),
		},
		{
			desc:     "2011-10-10T14:48:00, pattern4 1",
			input:    "2011-10-10T14:48:00",
			expected: time.Date(2011, time.Month(10), 10, 14, 48, 0, 0, getTimeZone("UTC")),
		},
		{
			desc:     "2011-10-10T14:48:00 GMT, pattern4 2",
			input:    "2011-10-10T14:48:00 GMT",
			expected: time.Date(2011, time.Month(10), 10, 14, 48, 0, 0, getTimeZone("GMT")),
		},
		{
			desc:     "2011-10-10T14:48:00.000+09:00, pattern5",
			input:    "2011-10-10T14:48:00.000+09:00",
			expected: time.Date(2011, time.Month(10), 10, 14, 48, 0, 0, getTimeZoneInt(9, 0)),
		},
		{
			desc:     "January 1, 1970 00:00:00 UTC, pattern6 2",
			input:    "January 1, 1970 00:00:00 UTC",
			expected: time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, getTimeZone("UTC")),
		},
		{
			desc:     "January 1, 09 00:00:00.00 GMT, pattern6 2",
			input:    "January 1, 09 00:00:00.00 GMT",
			expected: time.Date(2009, time.Month(1), 1, 0, 0, 0, 0, getTimeZone("GMT")),
		},
		{
			desc:     "January 1, 1970 00:00:00, pattern6 3",
			input:    "January 1, 1970 00:00:00",
			expected: time.Date(1970, time.Month(1), 1, 0, 0, 0, 0, getTimeZone("UTC")),
		},
		{
			desc:     "August 7, 2014, pattern6 4",
			input:    "August 7, 2014",
			expected: time.Date(2014, time.Month(8), 7, 0, 0, 0, 0, getTimeZone("UTC")),
		},
		{
			desc:     "Wed, 09 Aug 1995 00:00:00 GMT, pattern7 1",
			input:    "Wed, 09 Aug 1995 00:00:00 GMT",
			expected: time.Date(1995, time.Month(8), 9, 0, 0, 0, 0, getTimeZone("GMT")),
		},
		{
			desc:     "Wed, 09 Aug 1995 00:00:00, pattern7 2",
			input:    "Wed, 09 Aug 1995 00:00:00",
			expected: time.Date(1995, time.Month(8), 9, 0, 0, 0, 0, getTimeZone("UTC")),
		},
		{
			desc:     "Mon Jan 2 15:04:05 MST 2006, pattern8",
			input:    "Mon Jan 2 15:04:05 MST 2006",
			expected: time.Date(2006, time.Month(1), 2, 15, 4, 5, 0, getTimeZone("MST")),
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := gotime.Parse(tC.input)
			if err != nil {
				t.Errorf("TestError: %v", err)
			}
			if got.String() != tC.expected.String() {
				t.Errorf("TestFailed: got : %v != expected : %v", got, tC.expected)
			}
		})
	}
}
