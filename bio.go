package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

var (
	bioEN = []string{
		`Hi!`,
		``,
		`I am David Gageot`,
		`Software Engineer @Docker in Paris`,
		`I've been ❤ing Java for the past 20 years`,
		`No I do Go for the BETTER and the worse`,
		`I'm addicted to ultra-fast feedback.`,
		``,
		`You might know my blog JavaBien.net`,
	}

	bioFR = []string{
		`Hi!`,
		``,
		`I am David Gageot`,
		`Software Engineer @Docker in Paris`,
		`I've been ❤ing Java for the past 20 years`,
		`No I do Go for the BETTER and the worse`,
		`I'm addicted to ultra-fast feedback.`,
		``,
		`You might know my blog JavaBien.net`,
	}
)

func main() {
	lines := bioFR
	if len(os.Args) > 1 && os.Args[1] == "EN" {
		lines = bioEN
	}

	bio := centerLines(lines)

	fmt.Println(bio)
}

func centerLines(lines []string) string {
	text := ""

	longestLineLength := 0
	for _, line := range lines {
		lineLength := utf8.RuneCountInString(line)

		if lineLength > longestLineLength {
			longestLineLength = lineLength
		}
	}

	for _, line := range lines {
		lineLength := utf8.RuneCountInString(line)
		paddingLeft := (longestLineLength - lineLength) / 2

		formatString := fmt.Sprintf("%%%ds\n", lineLength+paddingLeft)

		text += fmt.Sprintf(formatString, line)
	}

	return text
}
