package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result int
	counter := 0
	for _, slice := range tab {
		f(slice)
		if slice >= "0" && slice <= "9" {
			counter++
		}
	}
	return result + counter
}
