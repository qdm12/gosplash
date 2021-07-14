package title

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getLeftRightCounts(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		lineLength int
		middle     string
		left       int
		right      int
	}{
		"no left or right": {
			lineLength: 12,
			middle:     "hello worlds",
		},
		"middle is too long": {
			lineLength: 11,
			middle:     "hello worlds",
		},
		"not enough space for left and right": {
			lineLength: 13,
			middle:     "hello worlds",
		},
		"one left and one right": {
			lineLength: 14,
			middle:     "hello worlds",
			left:       1,
			right:      1,
		},
		"one left and two right": {
			lineLength: 15,
			middle:     "hello worlds",
			left:       1,
			right:      2,
		},
		"two left and two right": {
			lineLength: 16,
			middle:     "hello worlds",
			left:       2,
			right:      2,
		},
		"utf-8 special character": {
			lineLength: 15,
			middle:     "❤️ hello ❤️",
			//          "x❤️ hello ❤️xx"
			//          "xxxxxxxxxxxxxxx" - 15 characters
			left:  1,
			right: 2,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			left, right := getLeftRightCounts(testCase.lineLength, testCase.middle)

			assert.Equal(t, testCase.left, left)
			assert.Equal(t, testCase.right, right)
		})
	}
}

func Test_surroundByRunes(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		middle string
		r      rune
		left   int
		right  int
		s      string
	}{
		"no left or right": {
			middle: "hello world",
			r:      '=',
			s:      "hello world",
		},
		"left and right": {
			middle: "hello world",
			r:      '=',
			left:   2,
			right:  3,
			s:      "==hello world===",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			s := surroundByRunes(testCase.middle, testCase.r, testCase.left, testCase.right)

			assert.Equal(t, testCase.s, s)
		})
	}
}
