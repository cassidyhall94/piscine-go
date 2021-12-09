package piscine

// Write a function that returns true if the string passed as a parameter contains only numerical characters, otherwise, returns false.

func IsNumeric(s string) bool {
	counter := 0
	for _, stringrune := range []rune(s) {
		if (stringrune >= rune(48)) && (stringrune <= rune(57)) {
			counter++
		} else {
			return false
		}
	}
	return true
}
