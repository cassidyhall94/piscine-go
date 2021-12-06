package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 200 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		result := nb
		for i := nb; i > 1; i-- {
			result = result * (i - 1)
		}
		return result
	}
}
