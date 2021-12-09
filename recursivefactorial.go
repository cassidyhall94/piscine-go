package piscine

// Write a recursive function that returns the factorial of the int passed as parameter.
// Errors (non possible values or overflows) will return 0.
// for is forbidden for this exercise.

func RecursiveFactorial(nb int) int {
	result := nb
	if nb < 0 || nb > 200 { // number range less than 0 and greater than 200
		return 0
	} else if nb == 0 { // number equals 0 then return 1
		return 1
	} else {
		if nb <= 1 { // if number is less than or equal to 1
			return 1
		}
		if nb > 1 { // if number is greater than 1
			return nb * RecursiveFactorial(nb-1) // multiply the number by the function again (number - 1)
		}
		return result
	}
}
