package piscine

func IsPrintable(s string) bool {
	counter := 0
	for _, stringrune := range []rune(s) {
		if (stringrune >= rune(32)) && (stringrune <= rune(126)) {
			counter++
		} else {
			return false
		}
	}
	return true
}
