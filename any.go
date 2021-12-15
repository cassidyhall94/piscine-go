package piscine

func Any(f func(string) bool, a []string) bool {
	for _, slice := range a {
		if slice >= "0" && slice <= "9" {
			f(slice)
			return true
		}
	}
	return false
}
