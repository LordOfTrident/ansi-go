package ansi

import "regexp"

// Regex from https://github.com/acarl005/stripansi
var Regex = regexp.MustCompile("[\u001b\u009b][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z" +
                               "\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf" +
                               "-ntqry=><~]))")
func Strip(text string) string {
	return Regex.ReplaceAllString(text, "")
}

func FindNext(text string) []int {
	return Regex.FindStringIndex(text)
}

func Count(text string) int {
	return len(Regex.FindAllString(text, -1))
}

func LengthNoAnsi(text string) int {
	return len(text) - Count(text)
}

func Contains(text string) bool {
	return Regex.MatchString(text)
}
