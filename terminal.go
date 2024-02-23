package ansi

import "fmt"

const (
	ResetDevice = "\x1bc"
)

func LineWrap(enable bool) string {
	return "\x1b[7" + ansiBoolFlag(enable)
}

func Title(title string) string {
	return fmt.Sprintf("\x1b[0;%v\x07", title)
}
