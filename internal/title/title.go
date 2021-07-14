package title

import "strings"

func MakeLines(separator rune, lineLength int,
	repository, madeByPrefix string, authors []string) (lines []string) {
	separatorLine := strings.Repeat(string(separator), lineLength)

	lines = append(lines, separatorLine, separatorLine)

	if repository != "" {
		projectNameWithSpaces := " " + repository + " "
		left, right := getLeftRightCounts(lineLength, projectNameWithSpaces)
		projectNameLine := surroundByRunes(projectNameWithSpaces, separator, left, right)
		lines = append(lines, projectNameLine, separatorLine)
	}

	madeByLines := makeMadeBy(madeByPrefix, authors, lineLength, separator)

	if len(madeByLines) > 0 {
		lines = append(lines, madeByLines...)
		lines = append(lines, separatorLine)
	}

	lines = append(lines, separatorLine)
	return lines
}
