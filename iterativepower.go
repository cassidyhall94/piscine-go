package piscine

// Write an iterative function that returns the value of nb to the power of power.
// Negative powers will return 0. Overflows do not have to be dealt with.

func IterativePower(nb int, power int) int {
	result := nb
	if power < 0 {
		return 0
	}
	if nb == 0 && power == 0 {
		return 1
	} else {
		for i := 0; i < power-1; i++ {
			result = result * nb
		}
	}
	return result
}
