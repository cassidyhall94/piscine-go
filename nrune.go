package piscine

// Write a function that returns the nth rune of a string. If not possible, it returns 0.

func NRune(s string, n int) rune {
	if n > (len([]rune(s))) || n <= 0 { // n is greater than the rune array [] length/ not possible because you cant return the rune if n is greater
		return rune(0)
	}
	return []rune(s)[n-1] // nth rune, if n is valid number (less than or equal to length of array)
}
