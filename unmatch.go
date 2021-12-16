package piscine

func Unmatch(a []int) int {
	for i, v := range a {
		var match bool
		for j, x := range a {
			if i == j {
				continue
			}
			if v == x {
				match = true
				break
			}
		}
		if match == false {
			return v
		}
	}
	return -1
}
