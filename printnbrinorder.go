package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var num []int
	for n != 0 {
		num = append(num, n%10)
		n /= 10
	}
	if len(num) == 0 {
		z01.PrintRune(rune(n) + '0')
	} else {
		sortinterger := len(num)
		for i := 0; i < sortinterger; i++ {
			for j := 0; j < (sortinterger - 1 - i); j++ {
				if num[j] > num[j+1] {
					num[j], num[j+1] = num[j+1], num[j]
				}
			}
		}
		for i := range num {
			z01.PrintRune(rune(num[i]) + '0')
		}
	}
}
