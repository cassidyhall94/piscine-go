package piscine

func IsLower(s string) bool {
	counter := 0
	for _, stringrune := range []rune(s) {
		if (stringrune >= rune(97)) && (stringrune <= rune(122)) {
			counter++
		} else {
			return false
		}
	}
	return true
}
