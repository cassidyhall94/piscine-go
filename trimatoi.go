package piscine

// Write a function that transforms numbers within a string, into an int.
// If the - sign is encountered before any number it should determine the sign of the returned int.
// This function should only return an int. In the case of an invalid input, the function should return 0.
// Note: There will never be more than one sign in a string in the tests.

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	for _, sb := range s {
		if sb == '-' && result == 0 {
			sign *= -1
		}
		if sb >= '0' && sb <= '9' {
			b := 0
			for i := 0; i < int(sb-'0'); i++ {
				b++
			}
			result = result*10 + b
		}
	}
	return result * sign
}
