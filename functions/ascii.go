package ascii

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Ascii(worr string) string {
	Filename := "standard.txt"
	file, err := os.Open(Filename)
	if err != nil {
		log.Fatal("Error opening file 😡 :", err)
	}
	defer file.Close()

	run := []string{}
	AsciiMap := map[rune][]string{}
	count := 0
	space := ' '
	Myscanner := bufio.NewScanner(file)
	totalLines := 0

	for Myscanner.Scan() {
		text := Myscanner.Text()
		totalLines++
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

	splitSlice := strings.Split(worr, "\n")

	if totalLines != 854 {
		log.Fatal("There is some issues with your file standard.txt")
	}

	var laste string
	if strings.Replace(worr, "\n", "", -1) == "" {
		for i := 0; i < strings.Count(worr, "\n"); i++ {
			laste += "\n"
		}
	}
	laste = PrintAscii(splitSlice, AsciiMap)

	return laste 
}
