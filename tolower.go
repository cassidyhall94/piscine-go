package piscine

func ToLower(s string) string {
	result := []rune(s)
	for i := range result {
		if result[i] >= rune(65) && result[i] <= rune(90) {
			result[i] = result[i] + 32
		}
	}
	return string(result)
}
