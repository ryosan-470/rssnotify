package rss

import (
	"fmt"
	"testing"
)

func TestRss2String(t *testing.T) {
	testCases := []struct {
		rss2     Rss2
		expected string
	}{
		{
			rss2: Rss2{},
			expected: `
Title: 
Link: 
Description: 
PubDate: 
Item: 
`,
		},
	}

	for _, testCase := range testCases {
		s := fmt.Sprintf("%s", testCase.expected)
		if s != testCase.expected {
			t.Errorf("got\n%v\nbut want\n%v", s, testCase.expected)
		}
	}
}

func TestItemToString(t *testing.T) {
	testCases := []struct {
		item     Item
		expected string
	}{
		{
			item: Item{},
			expected: `
[
	Title: 
	Link: 
	Description: 
	Content: 
	PubDate: 
	Comments: 
]`,
		},
	}

	for _, testCase := range testCases {
		s := fmt.Sprintf("%s", testCase.expected)
		if s != testCase.expected {
			t.Errorf("got\n%v\nbut want\n%v", s, testCase.expected)
		}
	}

}
