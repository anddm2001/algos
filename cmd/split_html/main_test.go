package main

import (
	"reflect"
	"testing"
)

func TestIsBlockTag(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"<div>", true},
		{"</div>", true},
		{"<p>", true},
		{"</p>", true},
		{"<h1>", true},
		{"<span>", false},
	}

	for _, test := range tests {
		got := isBlockTag(test.input)
		if got != test.want {
			t.Errorf("isBlockTag(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestSplitIntoPassages(t *testing.T) {
	tests := []struct {
		input string
		want  [][]byte
	}{
		{
			"<div>Hello world.</div> How are you? I am fine!<p>Thank you.</p>",
			[][]byte{
				[]byte("Hello world"),
				[]byte("How are you"),
				[]byte("I am fine"),
				[]byte("Thank you"),
			},
		},
		{
			"No tags here. Just text.",
			[][]byte{
				[]byte("No tags here"),
				[]byte("Just text"),
			},
		},
	}

	for _, test := range tests {
		got := splitIntoPassages(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("splitIntoPassages(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}
