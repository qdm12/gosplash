package annoucement

import (
	"time"
)

func MakeLines(message string, now, expiration time.Time) (lines []string) {
	if message == "" {
		return nil
	}
	// error covered by a unit test
	if !expiration.IsZero() && now.After(expiration) {
		return nil
	}
	return []string{"📣 " + message, ""}
}
