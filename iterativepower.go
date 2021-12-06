package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	} else {
		power := nb
		for i := 0; i < nb+1; i++ {
			power = power * i
		}
		return power
	}
}
