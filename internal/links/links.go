package links

import (
	"strings"
)

func MakeLines(rootURL, user, repo string, emails []string,
	paypalUser, githubSponsor string) (lines []string) {
	rootURL = strings.TrimSuffix(rootURL, "/")
	if rootURL != "" && user != "" && repo != "" {
		repoURL := rootURL + "/" + user + "/" + repo
		lines = append(lines, "🔧 Need help? "+repoURL+"/discussions/new")
		lines = append(lines, "🐛 Bug? "+repoURL+"/issues/new")
		lines = append(lines, "✨ New feature? "+repoURL+"/issues/new")
		lines = append(lines, "☕ Discussion? "+repoURL+"/discussions/new")
	}

	if len(emails) > 0 {
		lines = append(lines, "💻 Email? "+strings.Join(emails, ", "))
	}

	sponsorLine := makeSponsorLine(paypalUser, githubSponsor)
	if sponsorLine != "" {
		lines = append(lines, sponsorLine)
	}

	return lines
}

func makeSponsorLine(paypalUser, githubSponsor string) (line string) {
	if paypalUser == "" && githubSponsor == "" {
		return ""
	}

	line = "💰 Help me?"
	if paypalUser != "" {
		line += " https://www.paypal.me/" + paypalUser
	}
	if githubSponsor != "" {
		line += " https://github.com/sponsors/" + githubSponsor
	}
	return line
}
