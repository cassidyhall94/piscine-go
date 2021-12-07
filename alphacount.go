package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, stringrune := range []rune(s) {
		if ((stringrune >= rune(65)) && (stringrune <= rune(90))) || ((stringrune >= rune(97)) && (stringrune <= rune(122))) {
			counter++ // means counter = counter + 1
		}
	}
	return counter
}
