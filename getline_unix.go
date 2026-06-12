//go:build !windows && !js

package bubbline

import (
	"os"

	"golang.org/x/sys/unix"
)

var stopSignals = []os.Signal{unix.SIGINT, unix.SIGTERM}
