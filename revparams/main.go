package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argue := os.Args
	for i := (len(argue) - 1); i > 0; i-- {
		for _, prog := range argue[i] {
			z01.PrintRune(prog)
		}
		z01.PrintRune('\n')
	}
}
