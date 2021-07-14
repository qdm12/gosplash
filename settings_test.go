package gosplash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Settings_setDefaults(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		initial  Settings
		expected Settings
	}{
		"empty settings": {
			expected: Settings{
				LineLength:   40,
				Separator:    '=',
				RootURL:      "https://github.com",
				MadeByPrefix: "Made with ❤️  by ",
				Version:      "unknown",
				Commit:       "unknown",
				BuildDate:    "unknown date",
			},
		},
		"default author to user": {
			initial: Settings{
				User: "user",
			},
			expected: Settings{
				LineLength:   40,
				Separator:    '=',
				RootURL:      "https://github.com",
				User:         "user",
				Authors:      []string{"https://github.com/user"},
				MadeByPrefix: "Made with ❤️  by ",
				Version:      "unknown",
				Commit:       "unknown",
				BuildDate:    "unknown date",
			},
		},
		"no defaults": {
			initial: Settings{
				LineLength:   10,
				Separator:    'x',
				RootURL:      "https://gitea.com",
				MadeByPrefix: "Made by ",
				Version:      "v2",
				Commit:       "8942342",
				BuildDate:    "2021-07-14",
			},
			expected: Settings{
				LineLength:   10,
				Separator:    'x',
				RootURL:      "https://gitea.com",
				MadeByPrefix: "Made by ",
				Version:      "v2",
				Commit:       "8942342",
				BuildDate:    "2021-07-14",
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testCase.initial.setDefaults()

			assert.Equal(t, testCase.expected, testCase.initial)
		})
	}
}
