package piscine

func IsAlpha(s string) bool {
	counter := 0
	for _, stringrune := range []rune(s) {
		if ((stringrune >= rune(65)) && (stringrune <= rune(90))) || ((stringrune >= rune(97)) && (stringrune <= rune(122))) || stringrune == rune(32) || (stringrune >= rune(48)) && (stringrune <= rune(57)) {
			counter++
		} else {
			return false
		}
	}
	return true
}
