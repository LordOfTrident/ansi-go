package main

import (
	"time"
	"strings"
	"os"
	"fmt"

	"golang.org/x/term"

	"github.com/LordOfTrident/ansi-go"
)

var pink = ansi.FgHex("F27BCC")

func Pause(secs int, text string) {
	ansi.Println(ansi.Cursor(false),
	             ansi.Bold, ansi.LightWhite, text)

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

	ansi.Println(ansi.Underline, ansi.Bold, pink, "Hello, world!", ansi.CursorDown(1))

	Pause(5, "Pausing for 5 second(s)!")

	ansi.Print(ansi.RestoreScreen,
	           ansi.RestoreCursor)
}
