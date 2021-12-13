package piscine

func SplitWhiteSpaces(s string) []string {
	var printstr []string
	j := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == ' ' || rune(s[i]) == '\n' {
			if i != j {
				printstr = append(printstr, s[j:i])
			}
			j = i + 1
		} else if i == len(s)-1 {
			printstr = append(printstr, s[j:])
		}
	}
	return printstr
}
