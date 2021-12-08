package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argue := os.Args
	for _, prog := range argue[0] {
		z01.PrintRune(prog)
	}
	z01.PrintRune('\n')
}
