package piscine

func RecursiveFactorial(nb int) int {
	result := nb
	if nb <= 1 {
		return 1
	}
	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return result
}
