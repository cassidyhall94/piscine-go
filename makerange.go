package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var array []int
		return array
	} else {
		var result []int = make([]int, min, max)
		for i := min; i < max; i++ {
			result[i] = i
		}
		return result
	}
}
