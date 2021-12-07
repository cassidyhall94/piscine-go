package piscine

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
