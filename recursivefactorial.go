package piscine

func RecursiveFactorial(nb int) int {
	result := nb
	if nb < 0 || nb > 200 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		if nb <= 1 {
			return 1
		}
		if nb > 1 {
			return nb * RecursiveFactorial(nb-1)
		}
		return result
	}
}
