package piscine

var runerange = []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

func Atoi(s string) int {
	result := 0
	minusy := false
	for i, sb := range s {
		if i == 0 {
			if sb == 43 {
				continue
			} else if sb == 45 {
				minusy = true
			}
		}
		for _, r := range runerange {
			if r == sb {
				b := 0
				for i := 0; i < int(sb-'0'); i++ {
					b++
				}
				result = result*10 + b
			} else {
				return 0
			}
		}
	}
	if minusy {
		result *= 1
	}
	return result
}
