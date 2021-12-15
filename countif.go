package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result int
	counter := 0
	for _, slice := range tab {
		if f(slice) {
			counter++
		}
	}
	return result + counter
}
