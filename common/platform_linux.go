package common

import (
	"os"
	"path/filepath"
	"time"
)

const (
	defaultRobocorpLocation = "$HOME/.robocorp"
	defaultHoloLocation     = "/opt/robocorp/ht"

	defaultSema4Location     = "$HOME/.sema4ai"
	defaultSema4HoloLocation = "/opt/sema4ai/ht"
)

func ExpandPath(entry string) string {
	intermediate := os.ExpandEnv(entry)
	result, err := filepath.Abs(intermediate)
	if err != nil {
		return intermediate
	}
	return result
}

func PlatformSyncDelay() {
	time.Sleep(3 * time.Millisecond)
}
