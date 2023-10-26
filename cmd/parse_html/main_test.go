package main

import (
	"strings"
	"testing"
)

func TestRemoveHeadings(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "<div><h1>Header 1</h1><p>Some text</p></div>",
			expected: "<div><p>Some text</p></div>",
		},
		{
			input:    "<h2>Header 2</h2>",
			expected: "",
		},
		{
			input:    "<h3>Header 3</h3><h4>Header 4</h4><h5>Header 5</h5>",
			expected: "",
		},
		{
			input:    "<div><h6>Header 6</h6><p>Some text</p></div>",
			expected: "<div><p>Some text</p></div>",
		},
		{
			input:    "<p>Some text</p>",
			expected: "<p>Some text</p>",
		},
	}

	for _, test := range tests {
		output := removeHeadings(test.input)
		if output != test.expected {
			t.Errorf("Test failed for input %v, expected %v, but got %v", test.input, test.expected, output)
		}
	}
}

func TestRandStringRunes(t *testing.T) {
	// Тестируем, что функция создает строку заданной длины
	length := 10
	result := randStringRunes(length)
	if len(result) != length {
		t.Errorf("Ожидается строка длиной %d, но получено %d", length, len(result))
	}

	// Тестируем, что функция создает случайные строки (повторно вызванная функция не должна возвращать то же значение)
	result2 := randStringRunes(length)
	if result == result2 {
		t.Errorf("Дважды вызванная функция вернула одно и то же значение: %s", result)
	}
}

func TestRandStringRunesEmpty(t *testing.T) {
	// Тестируем, что функция возвращает пустую строку при нулевой длине
	result := randStringRunes(0)
	if result != "" {
		t.Errorf("Ожидается пустая строка, но получено %s", result)
	}
}

func TestRandStringRunesNegativeLength(t *testing.T) {
	// Тестируем, что функция возвращает пустую строку при отрицательной длине
	result := randStringRunes(-5)
	if result != "" {
		t.Errorf("Ожидается пустая строка, но получено %s", result)
	}
}

func TestRandStringRunesAlphabet(t *testing.T) {
	// Тестируем, что функция создает строку, содержащую только символы из алфавита
	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := randStringRunes(10)
	for _, char := range result {
		if !strings.ContainsRune(allowedChars, char) {
			t.Errorf("Строка содержит запрещенные символы: %s", result)
		}
	}
}
