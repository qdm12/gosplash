package title

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeLines(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		separator    rune
		lineLength   int
		repository   string
		madeByPrefix string
		authors      []string
		lines        []string
	}{
		"title with repo and authors": {
			separator:    '=',
			lineLength:   40,
			repository:   "unit test",
			madeByPrefix: "made by ",
			authors: []string{
				"author1",
				"author2",
				"author3",
			},
			lines: []string{
				"========================================",
				"========================================",
				"============== unit test ===============",
				"========================================",
				"===== made by author1, author2 and =====",
				"=============== author3 ================",
				"========================================",
				"========================================",
			},
		},
		"title without repo": {
			separator:    '=',
			lineLength:   40,
			madeByPrefix: "made by ",
			authors: []string{
				"author1",
				"author2",
				"author3",
			},
			lines: []string{
				"========================================",
				"========================================",
				"===== made by author1, author2 and =====",
				"=============== author3 ================",
				"========================================",
				"========================================",
			},
		},
		"title without authors": {
			separator:    '=',
			lineLength:   40,
			repository:   "unit test",
			madeByPrefix: "made by ",
			authors:      []string{},
			lines: []string{
				"========================================",
				"========================================",
				"============== unit test ===============",
				"========================================",
				"========================================",
			},
		},
		"emojis": {
			separator:    '=',
			lineLength:   40,
			repository:   "unit ❤️ test",
			madeByPrefix: "❤️ made ❤️ by ❤️ ",
			authors: []string{
				"author1",
				"author2",
				"author3",
			},
			lines: []string{
				"========================================",
				"========================================",
				"============= unit ❤️ test =============",
				"========================================",
				"====== ❤️ made ❤️ by ❤️ author1, ======",
				"========= author2 and author3 ==========",
				"========================================",
				"========================================",
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			lines := MakeLines(testCase.separator, testCase.lineLength,
				testCase.repository, testCase.madeByPrefix, testCase.authors)
			assert.Equal(t, testCase.lines, lines)
		})
	}
}
