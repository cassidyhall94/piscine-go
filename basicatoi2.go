package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, num := range s {
		conv := int(num) - 48

		if num <= 0 || num >= 0 {
		} else {
			result = (result * 10) + conv
		}
	}
	return result
}
