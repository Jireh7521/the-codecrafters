package main

import "strings"

func processLine(line string) (string, bool) {
	original := line

	line = strings.TrimSpace(line)

	if line == "" {
		return "", false
	}

	if strings.HasPrefix(line, "TODO:") {
		line = strings.Replace(line, "TODO:", "✦ ACTION:", 1)
	}

	if line == strings.ToUpper(line) {
		line = strings.Title(strings.ToLower(line))
	}

	if original == strings.ToLower(original) {
		line = strings.ToUpper(line)
	}

	if strings.Contains(line, "REVERSE") {
		line = Reverse(line)
	}

	return line, true
}
