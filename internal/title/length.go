package title

import (
	"unicode"
	"unicode/utf8"
)

func getVisualLength(s string) (length int) {
	var nonASCII int
	for s != "" {
		r, size := utf8.DecodeRuneInString(s)
		if r == utf8.RuneError {
			break
		} else if unicode.IsOneOf(
			[]*unicode.RangeTable{unicode.Variation_Selector}, r) {
			s = s[size:]
			continue // runes from the same emoji
		}

		if size == 1 {
			length++
		} else {
			nonASCII++
		}
		s = s[size:]
	}

	const nonASCIIVisualSize = 2.5
	length += int(float64(nonASCII) * nonASCIIVisualSize)

	return length
}
