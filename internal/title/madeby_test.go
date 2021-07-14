package title

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makeMadeBy(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		prefix    string
		authors   []string
		maxLength int
		separator rune
		lines     []string
	}{
		"no author": {},
		"single author": {
			prefix:    "wrote by ",
			maxLength: 15,
			separator: '=',
			authors:   []string{"quentin"},
			lines: []string{
				"== wrote by ===",
				"=== quentin ===",
			},
		},
		"3 authors": {
			prefix:    "wrote by ",
			maxLength: 30,
			separator: '=',
			authors:   []string{"quentin", "quentinou", "quintin"},
			lines: []string{
				"===== wrote by quentin, ======",
				"=== quentinou and quintin ====",
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			lines := makeMadeBy(testCase.prefix, testCase.authors,
				testCase.maxLength, testCase.separator)

			assert.Equal(t, testCase.lines, lines)
		})
	}
}
