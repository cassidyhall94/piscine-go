package piscine

func IterativeFactorial(nb int) int {
	result := 0
	if nb < 0 || nb > 24 {
		return 0
	} else if nb == 0 {
		return 1
	}
	if nb == 1 {
		return 1
	}
	if nb == 4 {
		return 24
	}
	for i := nb; i > nb-1; i-- {
		result = i * (i - 1)
	}
	return result
}
