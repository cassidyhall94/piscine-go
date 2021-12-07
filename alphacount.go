package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, stringrune := range []rune(s) { // looping through s string by slice
		if ((stringrune >= rune(65)) && (stringrune <= rune(90))) || ((stringrune >= rune(97)) && (stringrune <= rune(122))) { // checks to see if any of the runes aren't part of the alphabet using ascii table numbers for upper and lower case
			counter++ // means counter = counter + 1
		}
	}
	return counter
}
