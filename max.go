package piscine

func Max(a []int) int {
	var result int
	for _, v := range a {
		if v > result {
			result = v
		}
	}
	return result
}
