package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else {
		power := nb
		for i := 0; i < nb+1; i++ {
			power = power * i
		}
		return power
	}
}
