package title

import (
	"strings"
)

func getLeftRightCounts(lineLength int, middle string) (left, right int) {
	middleLength := getVisualLength(middle)
	if middleLength >= lineLength-1 {
		// note: we don't want only 1 left or only 1 right
		return 0, 0
	}

	leftAndRightCount := lineLength - middleLength
	left = leftAndRightCount / 2  //nolint:gomnd
	right = leftAndRightCount / 2 //nolint:gomnd
	if leftAndRightCount%2 != 0 {
		right++
	}

	return left, right
}

func surroundByRunes(middle string, r rune, left, right int) string {
	return strings.Repeat(string(r), left) +
		middle +
		strings.Repeat(string(r), right)
}
