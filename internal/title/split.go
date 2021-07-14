package title

import (
	"strings"
	"unicode/utf8"
)

func splitStringVertically(s string, maxLength int) (lines []string) {
	var currentLine string
	for _, newWord := range strings.Fields(s) {
		switch {
		case currentLine == "":
			currentLine = newWord
		case utf8.RuneCountInString(currentLine)+
			utf8.RuneCountInString(newWord) >= maxLength:
			lines = append(lines, currentLine)
			currentLine = newWord
		default:
			currentLine += " " + newWord
		}
	}

	lines = append(lines, currentLine)
	return lines
}
