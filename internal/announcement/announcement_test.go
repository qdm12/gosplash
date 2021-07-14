package annoucement

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_MakeLines(t *testing.T) {
	t.Parallel()

	const someUnix = 10000
	const someUnixAfter = 10100

	testCases := map[string]struct {
		message    string
		now        time.Time
		expiration time.Time
		lines      []string
	}{
		"no input": {},
		"message without expiration": {
			message: "hello world",
			lines:   []string{"📣 hello world", ""},
		},
		"expired message": {
			message:    "hello world",
			now:        time.Unix(someUnixAfter, 0),
			expiration: time.Unix(someUnix, 0),
		},
		"non expired message": {
			message:    "hello world",
			now:        time.Unix(someUnix, 0),
			expiration: time.Unix(someUnixAfter, 0),
			lines:      []string{"📣 hello world", ""},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			lines := MakeLines(testCase.message, testCase.now, testCase.expiration)

			assert.Equal(t, testCase.lines, lines)
		})
	}
}
