package piscine

func BasicAtoi2(s string) int {
	result := 0
	var num rune := []rune(s)
	for _, num := range s {
		conv := int(num) - 48

		if num[i] < '0' || num[i] > '9' {
		} else {
			result = (result * 10) + conv
		}
	}
	return result
}
