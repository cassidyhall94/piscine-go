package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		var array []int
		return array
	} else {
		var result []int
		for i := min; i < max; i++ {
			result = append(result, i)
		}
		return result
	}
}
