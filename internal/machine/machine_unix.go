//go:build !windows

package machine

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func GetInfo() (string, error) {
	var utsName unix.Utsname
	err := unix.Uname(&utsName)
	if err != nil {
		return "", err
	}
	os := parseData(utsName.Sysname[:])
	kernel := parseData(utsName.Release[:])
	arch := parseData(utsName.Machine[:])
	return fmt.Sprintf("%s %s (%s)", os, kernel, arch), nil
}

func parseData(data any) string {
	switch typedData := data.(type) {
	case []int8:
		buffer := make([]byte, 0, len(typedData))
		for _, b := range typedData {
			if b == 0 {
				break
			}
			buffer = append(buffer, byte(b))
		}
		return string(buffer)
	case []byte:
		buffer := make([]byte, 0, len(typedData))
		for _, b := range typedData {
			if b == 0 {
				break
			}
			buffer = append(buffer, b)
		}
		return string(buffer)
	default:
		panic(fmt.Sprintf("unexpected data type: %T", data))
	}
}
