package piscine

func Compact(ptr *[]string) int {
	counter := 0
	for _, str := range *ptr {
		if str != "" {
			counter++
		}
	}
	counter2 := 0
	ptrprint := make([]string, counter)
	for _, str := range *ptr {
		if str != "" {
			ptrprint[counter2] = str
			counter2++
		}
	}
	*ptr = ptrprint
	return counter
}
