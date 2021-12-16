package piscine

func Itoa(n int) string {
	var buf []byte
	var r []byte
	var next int
	var right int

	for {
		if n < 0 {
			n = -1 * n
			r = append(r, '-')
		}
		next = n / 10
		right = n - next*10
		n = next
		buf = append(buf, byte('0'+right))
		if n == 0 {
			break
		}
	}
	for j := 0; j < len(buf); j++ {
		r = append(r, buf[len(buf)-j-1])
	}
	return string(r)
}
