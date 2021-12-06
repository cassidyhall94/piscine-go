package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if power < 0 {
		return 0
	} else {
		for i := 0; i < power-1; i++ {
			result = result * nb
		}
	}
	return result
}
