package main

import (
	"github.com/01-edu/z01"
)

func isNegative(nb int) {
	if nb >= 0 {
		fmt.println("T")
		z01.PrintRune('\n')

	} else {
		fmt.Println("F")
		z01.PrintRune('\n')
	}
	}
}