package main

import (
	"log"
	"regexp"
	"strings"
)

// writeLogMarkdown accepts a string of any length, and appends the string to the log file.
func writeLogMarkdown(input string, format string) {
	logMarkdownPrefix := ""
	logMarkdownSuffix := ""
	logMarkdownEmptyLine := ""

	switch format {
	case "h1":
		logMarkdownPrefix = "## "
	case "h2":
		logMarkdownPrefix = "### "
	case "h3":
		logMarkdownPrefix = "#### "
	case "block":
		logMarkdownPrefix = "> "
		logMarkdownEmptyLine = "> "
	case "listitem":
		logMarkdownPrefix = "- "
	case "input":
		if strings.HasPrefix(input, "note ") {
			input = ""
		} else {
			logMarkdownPrefix = "`"
			logMarkdownSuffix = "`"
		}
	case "logentry":
		logMarkdownPrefix = "`"
		logMarkdownSuffix = "`"
	}

	stripped := stripTermFormatting(input)

	if len(logMarkdownPrefix+stripped+logMarkdownSuffix) > 0 {
		log.Println(logMarkdownPrefix + stripped + logMarkdownSuffix)
		log.Println(logMarkdownEmptyLine)
	}
}

// stringTermFormatting returns a string after removing termui formatting.
// Example:
// Input - [WOO](fg:green)
// Output - WOO
func stripTermFormatting(s string) string {
	// regex for '(fg:color)': \(fg(.*?)\)
	// regex for '[WOO]': \[(.*?)\]

	// strip the square brackets
	data := strings.ReplaceAll(s, "[", "")
	data = strings.ReplaceAll(data, "]", "")

	// strip '(fg:color)'
	var parRegex = regexp.MustCompile(`\(fg(.*?)\)`)
	st := parRegex.ReplaceAllString(data, "")

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
