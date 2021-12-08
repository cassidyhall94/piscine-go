package piscine

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
