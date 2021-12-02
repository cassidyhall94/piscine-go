package piscine

func BasicAtoi2(s string) int {
	result := 0
	b := 0
	check := true
	a := []rune(s)
	for _, num := range a {
		if byte(num) >= 48 && byte(num) <= 57 {
			for i := '0'; i < num; i++ {
				b++
			}
			result = (result * 10) + b
			b = 0
		} else {
			check = false
		}
		if check {
			return result
		} else {
			return 0
		}
	}
