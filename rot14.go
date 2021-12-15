package piscine

func Rot14(s string) string {
	result := []rune(s)
	for _, result := range s {
		if result >= (65) && result <= (77) || result >= (97) && result <= (109) {
			return (string(result + 14))
		}
		if result >= (110) && result <= (122) || result >= (78) && result <= (90) {
			return (string(result - 12))
		}
	}
	return string(result)
}
