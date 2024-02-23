package ansi

import "fmt"

const (
	SaveCursor        =  "\x1b[s"
	RestoreCursor     =  "\x1b[u"
	SaveCursorAttr    =  "\x1b7"
	RestoreCursorAttr =  "\x1b8"
	GetCursorPosition =  "\x1b[6n"
)

func Cursor(enable bool) string {
	return  "\x1b[?25" + ansiBoolFlag(enable)
}

func Goto(row int, col int) string {
	return  fmt.Sprintf("\x1b[%v;%vH", row, col)
}

func CursorUp(n int) string {
	if n == 0 {
		return ""
	} else {
		return  fmt.Sprintf("\x1b[%vA", n)
	}
}

func CursorDown(n int) string {
	if n == 0 {
		return ""
	} else {
		return  fmt.Sprintf("\x1b[%vB", n)
	}
}

func CursorLeft(n int) string {
	if n == 0 {
		return ""
	} else {
		return  fmt.Sprintf("\x1b[%vD", n)
	}
}

func CursorRight(n int) string {
	if n == 0 {
		return ""
	} else {
		return  fmt.Sprintf("\x1b[%vC", n)
	}
}
