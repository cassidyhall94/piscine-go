package piscine

func Index(s string, toFind string) int {
	// looping through tofind and s strings to compare their letters, and print their index if they are equal.
	for _, toFindrune := range []rune(toFind) {
		for index, srune := range []rune(s) {
			if toFindrune == srune {
				return index
			}
		}
		return -1 // if there's an invalid rune in toFind
	}
	return -1 // for formatting
}
