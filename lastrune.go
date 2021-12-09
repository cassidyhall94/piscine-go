package piscine

// Write a function that returns the last rune of a string.

func LastRune(s string) rune {
	return []rune(s)[len([]rune(s))-1]
}
