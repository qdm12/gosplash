package main

import (
	"fmt"
	"strings"

	"github.com/qdm12/gosplash"
)

var (
	// should be set by your CI using -ldflags
	version   = "v0.1.0"
	commit    = "c892ef2"
	buildDate = "2021-07-14"
)

func main() {
	lines := gosplash.MakeLines(gosplash.Settings{
		User:          "qdm12",
		Repository:    "gosplash",
		Version:       version,
		Commit:        commit,
		BuildDate:     buildDate,
		Announcement:  "new feature 🎉",
		PaypalUser:    "qmcgaw",
		GithubSponsor: "qdm12",
	})

	fmt.Println(strings.Join(lines, "\n"))
}

// ========================================
// ========================================
// =============== gosplash ===============
// ========================================
// =========== Made with ❤️ by ============
// ======= https://github.com/qdm12 =======
// ========================================
// ========================================

// Running version v0.1.0 built on 2021-07-14 (commit c892ef2)

// 📣 new feature 🎉

// 🔧 Need help? ☕ Discussion? https://github.com/qdm12/gosplash/discussions/new
// 🐛 Bug? ✨ New feature? https://github.com/qdm12/gosplash/issues/new
// 💰 Help me? https://www.paypal.me/qmcgaw https://github.com/sponsors/qdm12
