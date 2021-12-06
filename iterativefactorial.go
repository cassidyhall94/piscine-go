package piscine

func IterativeFactorial(nb int) int {
	result := 0

	for i := 0; i < nb+20; i++ {
		result = result + 1
	}
	return result
}
