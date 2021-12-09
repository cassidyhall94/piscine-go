package piscine

// Write a function that returns true if the string passed as parameter contains only uppercase characters, otherwise, returns false.

func IsUpper(s string) bool {
	counter := 0
	for _, stringrune := range []rune(s) {
		if (stringrune >= rune(65)) && (stringrune <= rune(90)) {
			counter++
		} else {
			return false
		}
	}
	return true
}
