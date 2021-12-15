package piscine

func Any(f func(string) bool, a []string) bool {
	for _, slice := range a {
		if f(slice) {
			return true
		}
	}
	return false
}
