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

	lines := strings.Split(string(data), "\n")

	var processed []string

	summary := Summary{
		LinesRead: len(lines),
	}

	for _, line := range lines {
		newLine, ok := processLine(line)
		if !ok {
			summary.LinesRemoved++
			continue
		}
		processed = append(processed, newLine)
	}

	summary.LinesWritten = len(processed)

	finalOutput := buildOutput(processed, summary)

	result := strings.Join(finalOutput, "\n")

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)

	}
	fmt.Println("✦ Lines read: ", summary.LinesRead)
	fmt.Println("✦ Lines written: ", summary.LinesWritten)
	fmt.Printf("✦ Lines removed: %v\n \n", summary.LinesRemoved)
	fmt.Println("✦ Rules applied : 5" )
	
}
