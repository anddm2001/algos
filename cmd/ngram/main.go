package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func isRepeatedWords(nGramWords []string) bool {
	for _, word := range nGramWords[1:] {
		if word != nGramWords[0] {
			return false
		}
	}
	return true
}

func nGrams(text string, n int) []string {
	cleaner := regexp.MustCompile(`[[:punct:]]`)
	text = cleaner.ReplaceAllString(text, "")
	text = strings.ToLower(text)

	words := regexp.MustCompile(`\s+`).Split(text, -1)

	uniqueNGrams := make(map[string]bool)
	nGramList := make([]string, 0)

	for i := 0; i <= len(words)-n; i++ {
		nGramWords := words[i : i+n]

		if n > 1 && isRepeatedWords(nGramWords) {
			continue
		}

		var builder strings.Builder
		for j, word := range nGramWords {
			if j > 0 {
				builder.WriteRune(' ')
			}
			builder.WriteString(word)
		}

		nGram := builder.String()
		if _, found := uniqueNGrams[nGram]; !found {
			uniqueNGrams[nGram] = true
			nGramList = append(nGramList, nGram)
		}
	}

	sort.Slice(nGramList, func(i, j int) bool {
		return strings.Split(nGramList[i], " ")[0] < strings.Split(nGramList[j], " ")[0]
	})

	return nGramList
}

func main() {
	text := "Hello world! Goodbye world! Hello there! World is great!"
	n := 2
	nGramsList := nGrams(text, n)
	for _, nGram := range nGramsList {
		fmt.Println(nGram)
	}
}
