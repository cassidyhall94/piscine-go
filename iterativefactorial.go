package piscine

import "github.com/01-edu/z01"

func IterativeFactorial(nb int) int {
	result := 0

	for i := 1; i < nb+1; i++ {
		result = result + i
	}
	if nb < 0 || nb > 24 {
		z01.PrintRune(48)
	}
	return result
}
