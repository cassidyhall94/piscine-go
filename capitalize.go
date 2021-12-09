package piscine

// Write a function that capitalizes the first letter of each word and lowercases the rest.
// A word is a sequence of alphanumeric characters.

func Capitalize(s string) string {
	result := []rune(s)
	if s != "" {
		if result[0] >= 97 && result[0] <= 122 {
			result[0] -= 32
		}
	}
	for i, stringrune := range s {
		if i != 0 {
			if IsAlpha(string(result[i-1])) {
				if stringrune >= 65 && stringrune <= 90 {
					result[i] += 32
				}
			}
			if !(IsAlpha(string(result[i-1]))) {
				if stringrune >= 97 && stringrune <= 122 {
					result[i] -= 32
				}
			}
		}
	}
	return string(result)
}
