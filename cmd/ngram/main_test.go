package main

import (
	"reflect"
	"testing"
)

func TestNGrams(t *testing.T) {
	tests := []struct {
		text     string
		n        int
		expected []string
	}{
		{
			text:     "Hello, world! Hello, people! Hello hello...",
			n:        2,
			expected: []string{"hello world", "hello people", "people hello", "world hello"},
		},
		{
			text:     "Testing single word.",
			n:        1,
			expected: []string{"single", "testing", "word"},
		},
		{
			text:     "",
			n:        2,
			expected: []string{},
		},
		{
			text:     "Same same same same",
			n:        2,
			expected: []string{},
		},
		{
			text:     "Same same same same",
			n:        1,
			expected: []string{"same"},
		},
	}

	for _, test := range tests {
		result := nGrams(test.text, test.n)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test failed for input %v, expected %v, but got %v", test.text, test.expected, result)
		}
	}
}
