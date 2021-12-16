package piscine

func Join(strs []string, sep string) string {
	concatenation := ""
	for i, result := range strs {
		concatenation = concatenation + result
		if i != len(strs)-1 {
			concatenation = concatenation + sep
		}
	}
	return concatenation
}
