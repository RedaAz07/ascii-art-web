package ascii

func PrintAscii(splitSlice []string, AsciiMap map[rune][]string) string {
	var result string
	for _, word := range splitSlice {
		if word != "" {
			for line := 0; line < 8; line++ {
				for _, char := range word {
					if ascii, exist := AsciiMap[char]; exist {
						result += ascii[line]
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}
