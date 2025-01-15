package ascii

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func Ascii(worr string, typee string , w http.ResponseWriter) string {
	var Filename string

	if typee == "standard" {
		Filename = "standard.txt"
	} else if typee == "shadow" {
		Filename = "shadow.txt"
	} else if typee == "thinkertoy" {
		Filename = "thinkertoy.txt"
	}
	file, err := os.Open(Filename)

	if  err !=  nil  {
		http.Error(w,"efezfef",400)
  return "grg"
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

	var laste string
	if strings.Replace(worr, "\n", "", -1) == "" {
		for i := 0; i < strings.Count(worr, "\n"); i++ {
			laste += "\n"
		}
	}
	laste = PrintAscii(splitSlice, AsciiMap)

	return laste
}
