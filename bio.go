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
		`I am David Gageot, freelance geek`,
		`I❤ Java, Tests, Cloud, Front Dev. c[_] rulez`,
		`I'm addicted to ultra-fast feedback. ie. I'm not patient (-.-)Zzz…`,
		``,
		`You might know my blog JavaBien.net`,
	}

	bioFR = []string{
		`Bonjour,`,
		``,
		`Je suis David Gageot, geek freelance`,
		`J'❤ Java, les Tests, le Cloud, le dev Front. c[_] rulez`,
		`Je suis accro au feedback rapide. ie. impatient (-.-)Zzz…`,
		`Vous suivez peut-etre mon blog JavaBien.net`,
		``,
		`@dgageot`,
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
