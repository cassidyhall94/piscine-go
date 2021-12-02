package piscine

func BasicAtoi(s string) int {
	runes := []rune(s)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	return result
}
