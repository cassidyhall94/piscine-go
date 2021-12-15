package piscine

func Abort(a, b, c, d, e int) int {
	median := []int{a, b, c, d, e}
	sortinterger := len(median)
	for i := 0; i < sortinterger; i++ {
		for j := 0; j < (sortinterger - 1 - i); j++ {
			if median[j] > median[j+1] {
				median[j], median[j+1] = median[j+1], median[j]
			}
		}
	}
	return median[2]
}
