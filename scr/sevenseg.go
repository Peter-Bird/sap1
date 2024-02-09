package scr

import "fmt"

var digitToUnicode = map[rune][]string{
	'0': {" _ ", "| |", "|_|"},
	'1': {"   ", "  |", "  |"},
	'2': {" _ ", " _|", "|_ "},
	'3': {" _ ", " _|", " _|"},
	'4': {"   ", "|_|", "  |"},
	'5': {" _ ", "|_ ", " _|"},
	'6': {"   ", "|_ ", "|_|"},
	'7': {" _ ", "  |", "  |"},
	'8': {" _ ", "|_|", "|_|"},
	'9': {" _ ", "|_|", " _|"},
}

func (display *Screen) Convert(number int) []string {
	numberStr := fmt.Sprintf("%03d", number)
	result := make([]string, 3)
	for _, digit := range numberStr {
		for i, line := range digitToUnicode[digit] {
			result[i] += line
		}
	}
	return result
}
