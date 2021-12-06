package piscine

func IterativeFactorial(nb int) int {
	result := 0
	if result < 0 || result > 24 {
		return 0
	}
	for i := nb; i > nb-1; i-- {
		result = i * (i - 1) * (i - 2)
	}
	return result
}
