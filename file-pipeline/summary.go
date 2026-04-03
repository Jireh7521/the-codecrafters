package main

import "fmt"

type Summary struct {
	LinesRead    int
	LinesWritten int
	LinesRemoved int
}

func buildOutput(lines []string, s Summary) []string {
	var output []string

	output = append(output, "Gopher's Sentinel Field Report - Processed")

	for i, line := range lines {
		output = append(output, fmt.Sprintf("%d. %s", i+1, line))
	}

	output = append(output, "")
	output = append(output, "Summary:")
	output = append(output, fmt.Sprintf("✦ Lines read    : %d", s.LinesRead))
	output = append(output, fmt.Sprintf("✦ Lines written : %d", s.LinesWritten))
	output = append(output, fmt.Sprintf("✦ Lines removed : %d", s.LinesRemoved))
	output = append(output, "✦ Rules applied : 5")

	return output
}
