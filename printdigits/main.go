package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var num string = "0123456789"
	for i := 0; i <= 9; i++ {
		z01.PrintRune(rune(num[i]))
	}
	z01.PrintRune('\n')
}
