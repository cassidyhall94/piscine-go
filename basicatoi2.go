package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, num := range s {
		conv := int(num) - 48

		if num[i] < 48 || num[i] > 57 {
		} else {
			result = (result * 10) + conv
		}
	}
	return result
}
