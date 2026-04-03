// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Samuel Jireh]
// Squad:  [Gophers]

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:
// Number of lines: 20
// Normal report lines
// Lines in ALL CAPS
// Lines in all lowercase
// Lines starting with TODO:
// Lines with extra leading/trailing spaces

// Transformation rules (in order):
// 1. Trim all leading and trailing whitespace
// 2. Replace TODO: with ✦ ACTION:
// 3. Convert ALL CAPS lines to Title Case
// 4. Convert all lowercase lines to uppercase
// 5. Reverse the words in any line that contains the word REVERSE

// Output format:
// Header: Yes, Exact Text: "Gopher's Sentinel Field Report - Processed"
// Line numbering format : "1."
// Summary block: yes
//  	Fields :
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
//
//
// Terminal summary fields:
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
// ═══════════════════════════════════════════

package main

import (
	"fmt"
	"os"
	"strings"
)

func ToLower(text string) string {
	return strings.ToLower(text)
}

func Title(text string) string {
	u := strings.ToLower(text)
	return strings.Title(u)
}

func Reverse(text string) string {
	words := strings.Fields(text)
	for i, w := range words {
		runes := []rune(w)
		for a := 0; a < len(runes)/2; a++ {
			b := len(runes) - 1 - a
			runes[a], runes[b] = runes[b], runes[a]
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// func isDigit(c byte) bool {
// 	return c >= '0' && c <= '9'
// }

// func isSpace(c byte) bool {
// 	return c == ' ' || c == '\n' || c == '\t'
// }

func isWord(s string) bool {
	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			return true
		}
	}
	return false
}

func applyModifier(tokens []string, cmd string, n int) []string {
	count := 0

	for i := len(tokens) - 1; i >= 0 && count < n; i-- {
		if !isWord(tokens[i]) {
			continue
		}

		switch cmd {
		case "up":
			tokens[i] = strings.ToUpper(tokens[i])
		case "low":
			tokens[i] = strings.ToLower(tokens[i])
		}

		count++
	}

	return tokens
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(data)

	tokens := ToLower(text)
	tokens = Title(tokens)
	tokens = Reverse(tokens)

	result := (tokens)

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	// fmt.Println(ToLower(inputFile))

}
