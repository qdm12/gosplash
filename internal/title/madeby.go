package title

func makeMadeBy(prefix string, authors []string, maxLength int, separator rune) (lines []string) {
	if len(authors) == 0 {
		return nil
	}

	s := prefix
	for i := range authors {
		switch {
		case i == 0:
			s += authors[i]
		case i < len(authors)-1:
			s += ", " + authors[i]
		default:
			s += " and " + authors[i]
		}
	}

	const maxContentLengthPercent = 0.75 // leave 25% for the separators
	maxContentLength := int(float64(maxLength) * maxContentLengthPercent)

	lines = splitStringVertically(s, maxContentLength)

	for i := range lines {
		lines[i] = " " + lines[i] + " "
		left, right := getLeftRightCounts(maxLength, lines[i])
		lines[i] = surroundByRunes(lines[i], separator, left, right)
	}

	return lines
}
