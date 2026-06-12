//go:build windows || js

package editline

var canSuspendProcess = false

func suspendProcess() {}
