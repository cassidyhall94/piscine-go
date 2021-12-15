package piscine

func Rot14(s string) string {
	result := []rune{}
	for _, srune := range s {
		if srune >= (65) && srune <= (76) || srune >= (97) && srune <= (109) {
			result = append(result, srune+14)
		} else if srune >= (110) && srune <= (122) || srune >= (77) && srune <= (90) {
			result = append(result, srune-12)
		} else {
			result = append(result, srune)
		}
	}
	return string(result)
}
