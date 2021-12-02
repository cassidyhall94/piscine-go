package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, num := range s {
		conv := int(num) - 48
		result = (result * 10) + conv
	}
	if num[i] < 48 || num[i] > 57 {

		return 0
	} else {
		result *= 10
		result += int(num[i]) - 48
	}
}
	return result
}