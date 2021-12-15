package piscine

func ForEach(f func(int), a []int) {
	for _, slice := range a {
		f(slice)
	}
}
