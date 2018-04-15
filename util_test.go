package main

import "testing"
import "time"

func TestToTime(t *testing.T) {
	testCases := []struct {
		input    string
		expected time.Time
	}{
		{
			input:    "Fri, 13 Apr 2018 12:34:56 UTC",
			expected: time.Date(2018, time.April, 13, 12, 34, 56, 0, time.UTC),
		},
	}

	for _, testCase := range testCases {
		ret := ToTime(testCase.input)
		if ret != testCase.expected {
			t.Errorf("\ngot  %v\nwant %v", ret, testCase.expected)
		}
	}
}

func TestIsUpdated(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	testCases := []struct {
		updated, now time.Time
		interval     int
		expected     bool
	}{
		{
			interval: 5,
			updated:  time.Date(2000, time.January, 1, 0, 0, 0, 0, loc),
			now:      time.Date(2000, time.January, 1, 0, 4, 59, 99, loc),
			expected: true,
		},
		{
			interval: 5,
			updated:  time.Date(2000, time.January, 1, 0, 0, 0, 0, loc),
			now:      time.Date(2000, time.January, 1, 0, 5, 0, 0, loc),
			expected: false,
		},
		{
			interval: 5,
			updated:  time.Date(2000, time.January, 1, 0, 0, 0, 0, loc),
			now:      time.Date(2000, time.January, 1, 0, 5, 0, 1, loc),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		actual := IsUpdated(testCase.interval, testCase.updated, testCase.now)
		if actual != testCase.expected {
			t.Errorf("\ngot %v\nwant %v", actual, testCase.expected)
		}
	}
}
