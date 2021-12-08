package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argue := os.Args
	for i, prog := range argue {
		if i != 0 {
			for _, prog := range prog {
				z01.PrintRune(prog)
			}
			z01.PrintRune('\n')
		}
	}
}
