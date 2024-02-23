//go:build windows
// +build windows
package ansi

import (
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	stdout := windows.Handle(os.Stdout.Fd())

	// Enable ansi escape sequences on Windows
	var origMode uint32
	windows.GetConsoleMode(stdout, &origMode)
	windows.SetConsoleMode(stdout, origMode | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
