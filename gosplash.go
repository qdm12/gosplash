package gosplash

import (
	"time"

	annoucement "github.com/qdm12/gosplash/internal/announcement"
	"github.com/qdm12/gosplash/internal/links"
	"github.com/qdm12/gosplash/internal/title"
)

var timeNow = time.Now //nolint:gochecknoglobals

// MakeLines returns the lines to log out.
func MakeLines(settings Settings) (lines []string) {
	settings.setDefaults()

	now := timeNow()

	titleLines := title.MakeLines(settings.Separator, settings.LineLength,
		settings.Repository, settings.MadeByPrefix, settings.Authors)
	versionString := "Running version " + settings.Version +
		" built on " + settings.Created + " (commit " + settings.Commit + ")"
	if settings.BuiltBy != "" {
		versionString += " by " + settings.BuiltBy
	}
	announcementLines := annoucement.MakeLines(settings.Announcement, now, settings.AnnounceExp)
	linksLines := links.MakeLines(settings.RootURL, settings.User, settings.Repository,
		settings.Emails, settings.PaypalUser, settings.GithubSponsor)

	lines = append(lines, titleLines...)
	lines = append(lines, "", versionString, "")
	lines = append(lines, announcementLines...)
	lines = append(lines, linksLines...)
	return lines
}
