package piscine

var counter int

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	if start == 1 {
		return counter
	}
	counter++
	if start%2 == 0 {
		return CollatzCountdown(start / 2)
	} else {
		return CollatzCountdown((start * 3) + 1)
	}
	return counter
}
