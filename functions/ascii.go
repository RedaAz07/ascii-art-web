package ascii

import (
	"bufio"
	"os"
	"strings"
)

func Ascii(word string, typee string) string {
	var Filename string

	if typee == "standard" {
		Filename = "files/standard.txt"
	} else if typee == "shadow" {
		Filename = "files/shadow.txt"
	} else if typee == "thinkertoy" {
		Filename = "files/thinkertoy.txt"
	}

	file, err := os.Open(Filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	run := []string{}
	AsciiMap := map[rune][]string{}
	count := 0
	space := ' '
	Myscanner := bufio.NewScanner(file)

	for Myscanner.Scan() {
		text := Myscanner.Text()
		if text != "" {
			run = append(run, text)
			count++
		}
		if count == 8 {
			AsciiMap[space] = run
			space++
			run = []string{}
			count = 0
		}
	}

	splitSlice := strings.Split(word, "\n")

	var laste string
	if strings.Replace(word, "\n", "", -1) == "" {
		for i := 0; i < strings.Count(word, "\n"); i++ {
			laste += "\n"
		}
	}
	laste = PrintAscii(splitSlice, AsciiMap)

	return laste
}
