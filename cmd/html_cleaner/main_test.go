package main

import (
	"testing"
)

func TestRemoveHTMLTags(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "<p>Simple text</p>",
			expected: "Simple text",
		},
		{
			input:    "  <div>  Text with   extra spaces  </div>  ",
			expected: "Text with extra spaces",
		},
		{
			input:    "<ul><li>Item 1</li><li>Item 2</li></ul>",
			expected: "Item 1 Item 2",
		},
		{
			input:    "<h1>Title</h1><p>Paragraph <b>with</b> different <a href='#'>tags</a>.</p>",
			expected: "Title Paragraph with different tags.",
		},
		{
			input:    "<div    attr='value'   >Test</div>",
			expected: "Test",
		},
		{
			input:    "Text without HTML tags",
			expected: "Text without HTML tags",
		},
	}

	for _, test := range tests {
		result := removeHTMLTags(test.input)
		if result != test.expected {
			t.Errorf("Expected '%s', but got '%s'", test.expected, result)
		}
	}
}

func TestRemoveHTMLTags_LongString(t *testing.T) {
	longInput := "<p>" + string(make([]byte, 2000)) + "</p>"
	expected := string(make([]byte, 2000))

	result := removeHTMLTagsOptimazed(longInput)
	if result != expected {
		t.Errorf("Testing with long string failed. Length of the resulting string: %d", len(result))
	}
}
