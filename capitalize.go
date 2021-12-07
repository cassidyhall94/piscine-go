package piscine

func Capitalize(s string) string {
	result := []rune(s)
	for i, stringrune := range result {
		if result[i] >= rune(48) && result[i] >= rune(57) {
		}
		if result[i] >= rune(65) && result[i] <= rune(90) {
			result[i] = stringrune + 32
		} else {
			if result[i] >= rune(97) && result[i] <= rune(122) {
				result[i] = stringrune + 32
			}
		}
	}
	return string(result)
}
