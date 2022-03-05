package gotime_test

import (
	"testing"
	"time"

	"github.com/emylincon/gotime"
)

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
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := gotime.Parse(tC.input)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if got != tC.expected {
				t.Errorf("Error: got : %v != expected : %v", got, tC.expected)
			}
		})
	}
}
