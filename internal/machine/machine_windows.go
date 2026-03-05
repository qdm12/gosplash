//go:build windows

package machine

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"
)

func GetInfo() (string, error) {
	version := windows.RtlGetVersion()
	return fmt.Sprintf("windows %d.%d.%d (%s)",
		version.MajorVersion, version.MinorVersion, version.BuildNumber, runtime.GOARCH), nil
}
