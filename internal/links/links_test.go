package links

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeLines(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		rootURL       string
		user          string
		repo          string
		emails        []string
		paypalUser    string
		githubSponsor string
		lines         []string
	}{
		"no input": {},
		"repo information": {
			rootURL: "https://github.com",
			user:    "qdm12",
			repo:    "gosplash",
			lines: []string{
				"🔧 Need help? ☕ Discussion? https://github.com/qdm12/gosplash/discussions/new",
				"🐛 Bug? ✨ New feature? https://github.com/qdm12/gosplash/issues/new",
			},
		},
		"emails": {
			emails: []string{"quentin.mcgaw@gmail.com", "a@a.com"},
			lines: []string{
				"💻 Email? quentin.mcgaw@gmail.com, a@a.com",
			},
		},
		"sponsor information": {
			paypalUser:    "qmcgaw",
			githubSponsor: "qdm12",
			lines: []string{
				"💰 Help me? https://www.paypal.me/qmcgaw https://github.com/sponsors/qdm12",
			},
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			lines := MakeLines(testCase.rootURL, testCase.user, testCase.repo,
				testCase.emails, testCase.paypalUser, testCase.githubSponsor)

			assert.Equal(t, testCase.lines, lines)
		})
	}
}
