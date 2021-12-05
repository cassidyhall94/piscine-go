package piscine

import "github.com/01-edu/z01"

// func drawLine(x, y int, top, mid, bot string) {
// 	topLeft := top[0]
// 	topMid := top[1]
// 	topRight := top[2]

// 	midLeft := top[0]
// 	midMid := top[1]
// 	midRight := top[2]

// 	topLeft := top[0]
// 	botMid := top[1]
// 	botRight := top[2]
// }

func QuadA(x, y int) {
	if x == 0 || y == 0 {
	} else {
		if x != 0 || y != 0 {
			if x > 0 && y > 0 {
				z01.PrintRune('o')
				for i := 1; i <= (x - 2); i++ {
					z01.PrintRune('-')
				}
				if x > 1 {
					z01.PrintRune('o')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('\n')
				}
				for j := 1; j <= (y - 2); j++ {
					z01.PrintRune('|')
					for k := 1; k <= (x - 2); k++ {
						z01.PrintRune(' ')
					}
					if x > 1 {
						z01.PrintRune('|')
						z01.PrintRune('\n')
					} else {
						z01.PrintRune('\n')
					}
				}
				if y > 1 {
					z01.PrintRune('o')
					for l := 1; l <= (x - 2); l++ {
						z01.PrintRune('-')
					}
					if x > 1 {
						z01.PrintRune('o')
						z01.PrintRune('\n')
					} else {
						z01.PrintRune('\n')
					}
				}
			} else {
				z01.PrintRune('o')
				z01.PrintRune('\n')
			}
		}
	}
}
