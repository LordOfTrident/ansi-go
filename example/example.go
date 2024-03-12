package main

import (
	"time"
	"strings"
	"os"
	"fmt"

	"golang.org/x/term"

	"github.com/LordOfTrident/ansi-go"
)

// 3 digit hex color literals
var red   = ansi.BgHex("F00")
var green = ansi.BgHex("0F0")
var blue  = ansi.BgHex("00F")

var pink = ansi.FgHex("F27BCC")
// Prefixing the hex color literal with # would also work:
//     var pink = ansi.FgHex("#F27BCC")

func Pause(secs int, text string) {
	ansi.Println(ansi.Cursor(false),
	             ansi.Bold, ansi.LightWhiteFg, text)

	prefix  := pink + "| " + ansi.Reset
	dotsLen := 4

	ansi.Println(prefix, ansi.Dim, "Please be patient :)")

	for i := 0; i < secs * dotsLen; i ++ {
		ansi.Print(ansi.ClearLine, "\r", prefix)
		ansi.Printf("%v more second(s)%v", secs - i / dotsLen, strings.Repeat(".", i % dotsLen))

		time.Sleep(time.Second / time.Duration(dotsLen))
	}

	ansi.Print(ansi.ClearLine,
	           ansi.CursorUp(1),
	           ansi.ClearLine,
	           "\r",
	           ansi.Cursor(true))
}

func main() {
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println("Oops! Looks like the output does not support ANSI escape sequences, shame :(")
		os.Exit(0)
	}

	ansi.Print(ansi.SaveCursor,
	           ansi.SaveScreen,
	           ansi.Goto(1, 1))

	textToSlice := ansi.RedFg + "Hello, " + ansi.GreenFg + "World!"
	ansi.Println("Text to slice: ", textToSlice)

	// ansi.InclusiveSlice includes the ansi sequences right before the first character, where as
	// ansi.Slice does not
	ansi.Println("Slicing with ansi: ", ansi.InclusiveSlice(textToSlice, 7, 12))

	ansi.Println(red, "  ", green, "  ", blue, "  ")
	ansi.Println(ansi.Underline, ansi.Bold, pink, "Hello, world!", ansi.CursorDown(1))

	Pause(5, "Pausing for 5 second(s)!")

	ansi.Print(ansi.RestoreScreen,
	           ansi.RestoreCursor)

	// Strip the sequences
	ansi.Println(ansi.Strip(ansi.Bold + ansi.BlueFg + "Goodbye." + ansi.Reset))
}
