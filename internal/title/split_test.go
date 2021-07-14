package title

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitStringVertically(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		s         string
		maxLength int
		lines     []string
	}{
		"empty string": {
			s:         "",
			maxLength: 10,
			lines:     []string{""},
		},
		"fit in one line": {
			s:         "hello world",
			maxLength: 11,
			lines:     []string{"hello world"},
		},
		"one long word": {
			s:         "hello",
			maxLength: 2,
			lines:     []string{"hello"},
		},
		"two long words": {
			s:         "hello world",
			maxLength: 2,
			lines:     []string{"hello", "world"},
		},
		"one word per line": {
			s:         "hello world how is it going it has been a very long time",
			maxLength: 5,
			lines:     []string{"hello", "world", "how", "is it", "going", "it", "has", "been", "a", "very", "long", "time"},
		},
		"multiple words per line": {
			s:         "hello world how is it going it has been a very long time",
			maxLength: 15,
			lines:     []string{"hello world how", "is it going it", "has been a very", "long time"},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			lines := splitStringVertically(testCase.s, testCase.maxLength)
			assert.Equal(t, testCase.lines, lines)
		})
	}
}
