package ansi

import "fmt"

const (
	ScrollDown         = "\x1bD"
	ScrollUp           = "\x1bM"
	ScrollScreen       = "\x1br"
	ClearScreen        = "\x1b[2J"
	ClearToScreenStart = "\x1b[1J"
	ClearToScreenEnd   = "\x1b[J"
	ClearLine          = "\x1b[2K"
	ClearToLineStart   = "\x1b[1K"
	ClearToLineEnd     = "\x1b[K"
	SaveScreen         = "\x1b[?47h"
	RestoreScreen      = "\x1b[?47l"
)

func AltScreen(enable bool) string {
	return "\x1b[?1049" + ansiBoolFlag(enable)
}

func ScrollViewport(start int, end int) string {
	return fmt.Sprintf("\x1b[%v;%vr", start, end)
}
