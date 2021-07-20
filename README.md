# GoSplash

Library to show log lines at the start of a program

[![Build status](https://github.com/qdm12/gosplash/actions/workflows/ci.yml/badge.svg)](https://github.com/qdm12/gosplash/actions/workflows/ci.yml)

![Last release](https://img.shields.io/github/release/qdm12/gosplash?label=Last%20release)
![GitHub last release date](https://img.shields.io/github/release-date/qdm12/gosplash?label=Last%20release%20date)
![Commits since release](https://img.shields.io/github/commits-since/qdm12/gosplash/latest?sort=semver)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/gosplash.svg)](https://github.com/qdm12/gosplash/commits/main)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/gosplash.svg)](https://github.com/qdm12/gosplash/graphs/contributors)
[![GitHub closed PRs](https://img.shields.io/github/issues-pr-closed/qdm12/gosplash.svg)](https://github.com/qdm12/gosplash/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/gosplash.svg)](https://github.com/qdm12/gosplash/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/qdm12/gosplash.svg)](https://github.com/qdm12/gosplash/issues?q=is%3Aissue+is%3Aclosed)

[![Lines of code](https://img.shields.io/tokei/lines/github/qdm12/gosplash)](https://github.com/qdm12/gosplash)
![Code size](https://img.shields.io/github/languages/code-size/qdm12/gosplash)
![GitHub repo size](https://img.shields.io/github/repo-size/qdm12/gosplash)
![Go version](https://img.shields.io/github/go-mod/go-version/qdm12/gosplash)

[![MIT](https://img.shields.io/github/license/qdm12/gosplash)](https://github.com/qdm12/gosplash/master/LICENSE)
![Visitors count](https://visitor-badge.laobi.icu/badge?page_id=gosplash.readme)

```log
==================================================
==================================================
==================== gosplash ====================
==================================================
================ Made with ❤️ by =================
============ https://github.com/qdm12 ============
==================================================
==================================================

Running version v1.1.1 built on 2021-07-14 (commit c892ef2)

📣 hello world

🔧 Need help? https://github.com/qdm12/gosplash/discussions/new
🐛 Bug? https://github.com/qdm12/gosplash/issues/new
✨ New feature? https://github.com/qdm12/gosplash/issues/new
☕ Discussion? https://github.com/qdm12/gosplash/discussions/new
💻 Email? quentin.mcgaw@gmail.com
💰 Help me? https://www.paypal.me/qmcgaw https://github.com/sponsors/qdm12
```

## Features

- Tailored towards Github repositories for now, create an issue if you need something
- Wide range of optional settings available:

    ```go
    type Settings struct {
      // Formatting settings
      Format       Format // defaults to DefaultFormat
      LineLength   int    // defaults to 40
      Separator    rune   // defaults to '='
      MadeByPrefix string // defaults to "Made with ❤️  by "

      // Project information
      RootURL    string // defaults to https://github.com
      User       string
      Repository string
      Authors    []string // defaults to [RootURL + "/" + User]
      Emails     []string

      // Program information
      Version      string // defaults to "unknown"
      Commit       string // defaults to "unknown"
      BuildDate    string // defaults to "unknown date"
      Announcement string
      AnnounceExp  time.Time // leave to zero value to disable expiration

      // Sponsor information
      PaypalUser    string
      GithubSponsor string
    }
    ```

- Support for emojis
- 100% test coverage

## Setup

```sh
go get github.com/qdm12/gosplash
```

From [the example](examples/main/main.go):

```go
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
```

## Quick links

- Problem or suggestion?
  - [Start a discussion](https://github.com/qdm12/gosplash/discussions)
  - [Create an issue](https://github.com/qdm12/gosplash/issues)
- Happy?
  - Sponsor me on [github.com/sponsors/qdm12](https://github.com/sponsors/qdm12)
  - Donate to [paypal.me/qmcgaw](https://www.paypal.me/qmcgaw)
  - Drop me [an email](mailto:quentin.mcgaw@gmail.com)
