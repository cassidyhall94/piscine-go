package piscine

// Write an iterative function that returns the factorial of the int passed as parameter.
// Errors (non possible values or overflows) will return 0.

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 200 { // number range less than 0 and greater than 200
		return 0
	} else if nb == 0 { // number equals 0 then return 1
		return 1
	} else {
		result := nb
		for i := nb; i > 1; i-- { // range through number and subject by one each time
			result = result * (i - 1) // result multiplied by i - 1
		}
		return result
	}
}
