package main

import (
	"regexp"
	"strings"
)

// blockTagsMap карта блочных HTML тегов для поиска и удаления
var blockTagsMap = map[string]bool{
	"<div>": true, "</div>": true,
	"<p>": true, "</p>": true,
	"<h1>": true, "</h1>": true,
	"<h2>": true, "</h2>": true,
	"<h3>": true, "</h3>": true,
}

// isBlockTag проверяет, является ли строка блочным HTML тегом
func isBlockTag(s string) bool {
	_, exists := blockTagsMap[s]
	return exists
}

// splitIntoPassages разбивает текст на пассажи, удаляя блочные теги и знаки препинания
func splitIntoPassages(text string) [][]byte {
	// Удаляем HTML теги и сохраняем содержимое тегов как отдельные пассажи
	for tag := range blockTagsMap {
		text = strings.ReplaceAll(text, tag, " ")
	}

	// Разбиваем текст на предложения по знакам препинания
	sentences := regexp.MustCompile(`[.!?]`).Split(text, -1)

	// Сохраняем каждый пассаж как срез байт в итоговом слайсе
	var passages [][]byte
	for _, sentence := range sentences {
		trimmed := strings.TrimSpace(sentence)
		if trimmed != "" {
			passages = append(passages, []byte(trimmed))
		}
	}

	return passages
}

func main() {
	text := "<div>Hello world.</div> How are you? I am fine!<p>Thank you.</p>"
	passages := splitIntoPassages(text)
	for _, passage := range passages {
		println(string(passage))
	}
}
