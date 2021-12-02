package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, num := range s {
		conv := int(num) - 48
		result = (result * 10) + conv
	}
	return result
}
