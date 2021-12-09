package piscine

// Write a function that capitalizes each letter in a string.

func ToUpper(s string) string {
	result := []rune(s)
	for i := range result {
		if result[i] >= rune(97) && result[i] <= rune(122) {
			result[i] = result[i] - 32
		}
	}
	return string(result)
}
