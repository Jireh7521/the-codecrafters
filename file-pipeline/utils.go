package main

import "strings"

func Title(text string) string {
	return strings.Title(strings.ToLower(text))
}

func Reverse(text string) string {
	words := strings.Fields(text)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
