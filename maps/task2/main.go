package main

import (
	"fmt"
	"strings"
	"unicode"
)

func WordFrequency(text string) map[string]int {
	result := make(map[string]int)
	fields := strings.FieldsFunc(strings.ToLower(text),
		func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsNumber(r) })
	for _, field := range fields {
		result[field]++
	}
	return result
}

func PrintWordFrequency(freq map[string]int) {
	for word, count := range freq {
		fmt.Printf("%s\t%d\n", word, count)
	}
}

func main() {
	text := "golang, is great and Golang is fast"
	freq := WordFrequency(text)
	PrintWordFrequency(freq)
}
