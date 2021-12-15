package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for _, slice := range a {
		result = append(result, f(slice))
	}
	return result
}
