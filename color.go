package ansi

import (
	"fmt"
	"strconv"
)

const (
	Black   = 0
	Red     = 1
	Green   = 2
	Yellow  = 3
	Blue    = 4
	Magenta = 5
	Cyan    = 6
	White   = 7

	LightBlack   = 8
	LightRed     = 9
	LightGreen   = 10
	LightYellow  = 11
	LightBlue    = 12
	LightMagenta = 13
	LightCyan    = 14
	LightWhite   = 15

	BlackFg   = "\x1b[30m"
	RedFg     = "\x1b[31m"
	GreenFg   = "\x1b[32m"
	YellowFg  = "\x1b[33m"
	BlueFg    = "\x1b[34m"
	MagentaFg = "\x1b[35m"
	CyanFg    = "\x1b[36m"
	WhiteFg   = "\x1b[37m"

	BlackBg   = "\x1b[40m"
	RedBg     = "\x1b[41m"
	GreenBg   = "\x1b[42m"
	YellowBg  = "\x1b[43m"
	BlueBg    = "\x1b[44m"
	MagentaBg = "\x1b[45m"
	CyanBg    = "\x1b[46m"
	WhiteBg   = "\x1b[47m"

	LightBlackFg   = "\x1b[90m"
	LightRedFg     = "\x1b[91m"
	LightGreenFg   = "\x1b[92m"
	LightYellowFg  = "\x1b[93m"
	LightBlueFg    = "\x1b[94m"
	LightMagentaFg = "\x1b[95m"
	LightCyanFg    = "\x1b[96m"
	LightWhiteFg   = "\x1b[97m"

	LightBlackBg   = "\x1b[100m"
	LightRedBg     = "\x1b[101m"
	LightGreenBg   = "\x1b[102m"
	LightYellowBg  = "\x1b[103m"
	LightBlueBg    = "\x1b[104m"
	LightMagentaBg = "\x1b[105m"
	LightCyanBg    = "\x1b[106m"
	LightWhiteBg   = "\x1b[107m"
)

func IsBright16(color uint8) bool {
	return color >= 8 && color <= 15
}

// TODO: If terminal does not support truecolor or 256 colors, find the closest supported color
func Fg(color uint8) string {
	return fmt.Sprintf("\x1b[38;5;%vm", color)
}

func Bg(color uint8) string {
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

	color32, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return
	}

	switch len(hex) {
	case 6:
		ok = true
		r  = uint8(color32 & 0xFF0000 >> 16)
		g  = uint8(color32 & 0x00FF00 >> 8)
		b  = uint8(color32 & 0x0000FF)

	case 3:
		ok = true
		r  = 17 * uint8(color32 & 0xF00 >> 8)
		g  = 17 * uint8(color32 & 0x0F0 >> 4)
		b  = 17 * uint8(color32 & 0x00F)
	}

	return
}
