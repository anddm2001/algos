package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Оптимизировано для больших строк
func removeHTMLTagsOptimazed(input string) string {
	reHeaders := regexp.MustCompile(`(<\/?h[1-6].*?>\s*)+`)
	input = reHeaders.ReplaceAllString(input, " ")

	re := regexp.MustCompile(`(?i)<.*?>`)
	reMultipleTags := regexp.MustCompile(`(<\/?\w+.*?>\s*){2,}`)

	input = reMultipleTags.ReplaceAllString(input, " ")
	cleanStr := re.ReplaceAllString(input, "")

	var b bytes.Buffer
	prevChar := ' '
	for _, c := range cleanStr {
		if c == ' ' && prevChar == ' ' {
			continue
		}
		b.WriteRune(c)
		prevChar = c
	}

	return string(bytes.TrimSpace(b.Bytes()))
}

func removeHTMLTags(input string) string {
	reHeaders := regexp.MustCompile(`(<\/?h[1-6].*?>\s*)+`)
	input = reHeaders.ReplaceAllString(input, " ")

	re := regexp.MustCompile(`(?i)<.*?>`)
	reMultipleTags := regexp.MustCompile(`(<\/?\w+.*?>\s*){2,}`)

	input = reMultipleTags.ReplaceAllString(input, " ")
	cleanStr := re.ReplaceAllString(input, "")
	cleanStr = strings.Join(strings.Fields(cleanStr), " ")

	return cleanStr
}

func main() {
	shortString := "<h1>Short String</h1>"
	longString := "<h1>" + string(make([]byte, 2000)) + "</h1>"

	// Тестирование с короткой строкой
	start := time.Now()
	removeHTMLTags(shortString)
	durationV1 := time.Since(start)

	start = time.Now()
	removeHTMLTagsOptimazed(shortString)
	durationV2 := time.Since(start)

	fmt.Printf("Duration for V1 with short string: %v\n", durationV1)
	fmt.Printf("Duration for V2 with short string: %v\n", durationV2)

	// Тестирование с длинной строкой
	start = time.Now()
	removeHTMLTags(longString)
	durationV1 = time.Since(start)

	start = time.Now()
	removeHTMLTagsOptimazed(longString)
	durationV2 = time.Since(start)

	fmt.Printf("Duration for V1 with long string: %v\n", durationV1)
	fmt.Printf("Duration for V2 with long string: %v\n", durationV2)
}
