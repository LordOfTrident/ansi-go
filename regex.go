package ansi

import (
	"regexp"
	"unicode/utf8"
)

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

func SizeNoAnsi(text string) int {
	return len(Strip(text))
}

func LengthNoAnsi(text string) int {
	return utf8.RuneCountInString(Strip(text))
}

func Contains(text string) bool {
	return Regex.MatchString(text)
}

// Convert rune index without ansi sequences to index including ansi sequences
func RunePosToByteIndex(text string, runePos int) int {
	it         := text
	pos        := 0
	snippetIdx := 0

	for {
		var snippet string

		loc := FindNext(it)
		if loc == nil {
			snippet = it
		} else {
			snippet = it[:loc[0]]
		}

		for i := range snippet {
			if pos == runePos {
				return snippetIdx + i
			}

			pos ++
		}

		if loc == nil {
			return -1
		}

		snippetIdx += loc[1]
		it          = it[loc[1]:]
	}
}

func Slice(text string, start int, end int) string {
	return text[RunePosToByteIndex(text, start):RunePosToByteIndex(text, end)]
}

// Include the ansi sequences before the first character
func InclusiveSlice(text string, start int, end int) string {
	end = RunePosToByteIndex(text, end)

	if start > 0 {
		start = RunePosToByteIndex(text, start - 1)
		return text[start:end][1:]
	} else {
		return text[start:end]
	}
}
