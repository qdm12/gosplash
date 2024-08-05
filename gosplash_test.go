package gosplash

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_MakeLines(t *testing.T) { //nolint:tparallel
	t.Parallel()

	const someUnix = 1000
	const someUnixAfter = 1100

	testCases := map[string]struct {
		now      time.Time
		settings Settings
		lines    []string
	}{
		"empty settings": {
			lines: []string{
				"========================================",
				"========================================",
				"========================================",
				"",
				"Running version unknown built on unknown date (commit unknown)",
				"",
			},
		},
		"filled settings": {
			now: time.Unix(someUnix, 0),
			settings: Settings{
				Format:        DefaultFormat,
				LineLength:    50,
				Separator:     '=',
				User:          "qdm12",
				Repository:    "gosplash",
				Emails:        []string{"quentin.mcgaw@gmail.com"},
				Version:       "v1.1.1",
				Commit:        "c892ef2",
				Created:       "2021-07-14",
				Announcement:  "hello world",
				AnnounceExp:   time.Unix(someUnixAfter, 0),
				PaypalUser:    "qmcgaw",
				GithubSponsor: "qdm12",
			},
			lines: []string{
				"==================================================",
				"==================================================",
				"==================== gosplash ====================",
				"==================================================",
				"================ Made with ❤️ by =================",
				"============ https://github.com/qdm12 ============",
				"==================================================",
				"==================================================",
				"",
				"Running version v1.1.1 built on 2021-07-14 (commit c892ef2)",
				"",
				"📣 hello world",
				"",
				"🔧 Need help? ☕ Discussion? https://github.com/qdm12/gosplash/discussions/new/choose",
				"🐛 Bug? ✨ New feature? https://github.com/qdm12/gosplash/issues/new/choose",
				"💻 Email? quentin.mcgaw@gmail.com",
				"💰 Help me? https://www.paypal.me/qmcgaw https://github.com/sponsors/qdm12",
			},
		},
	}

	for name, testCase := range testCases { //nolint:paralleltest
		t.Run(name, func(t *testing.T) {
			timeNow = func() time.Time {
				return testCase.now
			}
			defer func() {
				timeNow = time.Now
			}()

			lines := MakeLines(testCase.settings)

			assert.Equal(t, testCase.lines, lines)
		})
	}
}
