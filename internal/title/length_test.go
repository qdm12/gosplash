package title

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getVisualLength(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		s      string
		length int
	}{
		"empty": {},
		"bad rune": {
			s: string([]byte{255}),
		},
		"ascii characters": {
			s:      "abc",
			length: 3,
		},
		"emojis": {
			s: "ððð",
			// "aaaaaaa"
			length: 7,
		},
		"characters and emojis": {
			s: "a ðb ð c ðd",
			// "aaaaaaaaaaaaaaa"
			length: 15,
		},
		"long chain of characters and emojis": {
			s: "aaðaabðcðdðððvdasfðsðaaað",
			// "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			length: 38,
		},
		"heart emojis": {
			s: " â¤ï¸ made â¤ï¸ by â¤ï¸ author1, ",
			// "aaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			length: 28,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			length := getVisualLength(testCase.s)
			assert.Equal(t, testCase.length, length)
		})
	}
}
