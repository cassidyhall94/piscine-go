package piscine

func NRune(s string, n int) rune {
	if n > (len([]rune(s)) - 1) { // n is greater than the rune array [] length/ not possible because you cant return the rune if n is greater
		return []rune(s)[0] // first rune
	}
	return []rune(s)[n-1] // nth rune, if n is valid number (less than or equal to length of array)
}
