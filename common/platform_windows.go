package common

import (
	"os"
	"path/filepath"
	"regexp"
	"time"
)

const (
	defaultRobocorpLocation = "%LOCALAPPDATA%\\robocorp"
	defaultHoloLocation     = "%ProgramData%\\robocorp\\ht"

	defaultSema4Location     = "%LOCALAPPDATA%\\sema4ai"
	defaultSema4HoloLocation = "%ProgramData%\\sema4ai\\ht"
)

var (
	variablePattern = regexp.MustCompile("%[a-zA-Z]+%")
)

func ExpandPath(entry string) string {
	intermediate := os.ExpandEnv(entry)
	intermediate = variablePattern.ReplaceAllStringFunc(intermediate, fromEnvironment)
	result, err := filepath.Abs(intermediate)
	if err != nil {
		return intermediate
	}
	return result
}

func fromEnvironment(form string) string {
	replacement, ok := os.LookupEnv(form[1 : len(form)-1])
	if ok {
		return replacement
	}
	replacement, ok = os.LookupEnv(form)
	if ok {
		return replacement
	}
	return form
}

func PlatformSyncDelay() {
	time.Sleep(300 * time.Millisecond)
}
