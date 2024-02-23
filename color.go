package ansi

import (
	"fmt"
	"strconv"
)

const (
	Black   ="\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"

	BlackBg   = "\x1b[30m"
	RedBg     = "\x1b[31m"
	GreenBg   = "\x1b[32m"
	YellowBg  = "\x1b[33m"
	BlueBg    = "\x1b[34m"
	MagentaBg = "\x1b[35m"
	CyanBg    = "\x1b[36m"
	WhiteBg   = "\x1b[37m"

	LightBlack   = "\x1b[90m"
	LightRed     = "\x1b[91m"
	LightGreen   = "\x1b[92m"
	LightYellow  = "\x1b[93m"
	LightBlue    = "\x1b[94m"
	LightMagenta = "\x1b[95m"
	LightCyan    = "\x1b[96m"
	LightWhite   = "\x1b[97m"

	LightBlackBg   = "\x1b[90m"
	LightRedBg     = "\x1b[91m"
	LightGreenBg   = "\x1b[92m"
	LightYellowBg  = "\x1b[93m"
	LightBlueBg    = "\x1b[94m"
	LightMagentaBg = "\x1b[95m"
	LightCyanBg    = "\x1b[96m"
	LightWhiteBg   = "\x1b[97m"
)

func Fg256(color uint8) string {
	return fmt.Sprintf("\x1b[38;5;%vm", color)
}

func Bg256(color uint8) string {
	return fmt.Sprintf("\x1b[48;5;%vm", color)
}

func FgRGB(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%v;%v;%vm", r, g, b)
}

func BgRGB(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("\x1b[48;2;%v;%v;%vm", r, g, b)
}

func FgHex(hex string) string {
	r, g, b, ok := parseHex(hex)
	if !ok {
		return ""
	}

	return FgRGB(r, g, b)
}

func BgHex(hex string) string {
	r, g, b, ok := parseHex(hex)
	if !ok {
		return ""
	}

	return BgRGB(r, g, b)
}

func parseHex(hex string) (r uint8, g uint8, b uint8, ok bool) {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	if len(hex) != 6 {
		return
	}

	color32, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return
	}

	ok = true
	r  = uint8(color32 & 0xFF0000 >> 16)
	g  = uint8(color32 & 0x00FF00 >> 8)
	b  = uint8(color32 & 0x0000FF)
	return
}
