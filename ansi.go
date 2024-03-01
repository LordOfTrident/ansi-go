package ansi

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

const Marker = '\x1b'

func ansiBoolFlag(on bool) string {
	if on {
		return "h"
	} else {
		return "l"
	}
}

// Regex from https://github.com/acarl005/stripansi
var stripRegex = regexp.MustCompile("[\u001b\u009b][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z" +
                                    "\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf" +
                                    "-ntqry=><~]))")
func Strip(text string) string {
	return stripRegex.ReplaceAllString(text, "")
}

/*
 * All these print functions are almost the same as their fmt counterparts, the difference is
 * that they never add spaces between arguments and always add an ansi.Reset sequence at the end.
 */

func Fprintf(w io.Writer, format string, args ...any) (int, error) {
	return fmt.Fprintf(w, format + Reset, args...)
}

func Fprint(w io.Writer, args ...any) (int, error) {
	output := ""
	for _, arg := range args {
		output += fmt.Sprintf("%v", arg)
	}

	return fmt.Fprint(w, output, Reset)
}

func Fprintln(w io.Writer, args ...any) (int, error) {
	output := ""
	for _, arg := range args {
		output += fmt.Sprintf("%v", arg)
	}

	return fmt.Fprint(w, output + "\n", Reset)
}

// Stdout variants
func Printf(format string, args ...any) (int, error) {
	return Fprintf(os.Stdout, format, args...)
}

func Print(args ...any) (int, error) {
	return Fprint(os.Stdout, args...)
}

func Println(args ...any) (int, error) {
	return Fprintln(os.Stdout, args...)
}

// Stderr variants
func Eprintf(format string, args ...any) (int, error) {
	return Fprintf(os.Stderr, format, args...)
}

func Eprint(args ...any) (int, error) {
	return Fprint(os.Stderr, args...)
}

func Eprintln(args ...any) (int, error) {
	return Fprintln(os.Stderr, args...)
}
