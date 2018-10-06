package main

import (
	"testing"
)

var testCases = []struct {
	templ    string
	input    string
	expected int
}{
	{
		templ:    "12",
		input:    "12",
		expected: 0,
	},
	{
		templ:    "12",
		input:    "123",
		expected: 1,
	},
	{
		templ:    "12",
		input:    "11",
		expected: 1,
	},
	{
		templ:    "12",
		input:    "1",
		expected: 1,
	},
	{
		templ:    "12345",
		input:    "12",
		expected: 3,
	},
}

func TestDistance(t *testing.T) {
	for _, tt := range testCases {
		actual := Distance(tt.templ, tt.input)
		if actual != tt.expected {
			t.Fatalf("Distance(%v, %v) was expected to return %v but returned %d.",
				tt.templ, tt.input, tt.expected, actual)
		}
	}
}
