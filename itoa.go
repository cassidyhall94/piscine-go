package piscine

var runes = []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}

func Itoa(n int) string {
	var result string
	for i, sb := range result {
		if i == 0 {
			if sb == 43 {
				continue
			} else if sb == 45 {
				continue
			}
		}
		if sb < 58 && sb > 47 {
			b := 0
			for i := 0; i < int(sb-'0'); i++ {
				b++
			}
		} else {
			return result
		}
	}
	return result
}
