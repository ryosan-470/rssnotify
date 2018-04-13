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
