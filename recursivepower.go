package piscine

// Write a recursive function that returns the value of nb to the power of power.
// Negative powers will return 0. Overflows do not have to be dealt with.
// for is forbidden for this exercise.

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb == 0 && power == 0 {
		return 1
	} else if power != 0 {
		return nb * RecursivePower(nb, power-1)
	} else {
		return 1
	}
}
