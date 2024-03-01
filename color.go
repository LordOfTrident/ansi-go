package ansi

import (
	"fmt"
	"strconv"
)

const (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"

	BlackBg   = "\x1b[40m"
	RedBg     = "\x1b[41m"
	GreenBg   = "\x1b[42m"
	YellowBg  = "\x1b[43m"
	BlueBg    = "\x1b[44m"
	MagentaBg = "\x1b[45m"
	CyanBg    = "\x1b[46m"
	WhiteBg   = "\x1b[47m"

	LightBlack   = "\x1b[90m"
	LightRed     = "\x1b[91m"
	LightGreen   = "\x1b[92m"
	LightYellow  = "\x1b[93m"
	LightBlue    = "\x1b[94m"
	LightMagenta = "\x1b[95m"
	LightCyan    = "\x1b[96m"
	LightWhite   = "\x1b[97m"

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
	return color >= 8 && color <= 15;
}

func Fg16(color uint8) string {
	if IsBright16(color) {
		color += 82
	} else {
		color += 30
	}

	return fmt.Sprintf("\x1b[%vm", color)
}

func Bg16(color uint8) string {
	if IsBright16(color) {
		color += 92
	} else {
		color += 40
	}

	return fmt.Sprintf("\x1b[%vm", color)
}

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
