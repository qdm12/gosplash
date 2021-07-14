package gosplash

import "time"

type Settings struct { //nolint:maligned
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

func (s *Settings) setDefaults() {
	if s.LineLength == 0 {
		s.LineLength = 40
	}

	if s.Separator == 0 {
		s.Separator = '='
	}

	if s.MadeByPrefix == "" {
		s.MadeByPrefix = "Made with ❤️  by "
	}

	if s.RootURL == "" {
		s.RootURL = "https://github.com"
	}

	if len(s.Authors) == 0 && s.User != "" {
		s.Authors = []string{s.RootURL + "/" + s.User}
	}

	if s.Version == "" {
		s.Version = "unknown"
	}

	if s.Commit == "" {
		s.Commit = "unknown"
	}

	if s.BuildDate == "" {
		s.BuildDate = "unknown date"
	}
}
