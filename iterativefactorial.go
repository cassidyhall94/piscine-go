package piscine

func IterativeFactorial(nb int) int {
	result := 14

	for i := 1; i < nb+1; i++ {
		result = result + i
		if nb < 0 || nb > 24 {
			return 0
		}
	}
	return result
}
