package main

import (
	"log"
	"regexp"
	"strings"
)

// writeLogMarkdown accepts a string of any length, and appends the string to the log file.
func writeLogMarkdown(input string) {
	fullString := stripTermFormatting(input)
	words := strings.Fields(fullString)

	output := ""

	switch words[0] {
	case "---":
		output += "## " + GetStringInBetween(fullString, "--- ", " ---")
	case "--":
		output += "### " + GetStringInBetween(fullString, "-- ", " --")
	default:
		output += fullString
	}
	log.Println(output)
}

// stringTermFormatting returns a string after removing termui formatting.
// Example:
// Input - [WOO](fg:green)
// Output - WOO
func stripTermFormatting(s string) string {
	// regex for '(fg:color)': \(fg(.*?)\)
	// regex for '[WOO]': \[(.*?)\]

	// strip the square brackets
	test := strings.ReplaceAll(s, "[", "")
	test = strings.ReplaceAll(test, "]", "")

	// strip '(fg:color)'
	var parRegex = regexp.MustCompile(`\(fg(.*?)\)`)
	st := parRegex.ReplaceAllString(test, "")

	return st
}

// GetStringInBetween returns the contents between a start string and an end string.
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}
