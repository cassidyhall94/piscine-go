package piscine

// Write a function that behaves like the Compare function.

func Compare(a, b string) int {
	result := 0
	if a == b {
		return 0
	}
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return result
}
