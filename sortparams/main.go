package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argue := os.Args[1]
	for i, prog := range argue {
		for i, prog := range prog {
			for i := 0; i < argue; i++ {
				if prog[i] > prog[i+i] {
					prog[i], prog[i+1] = prog[i+1], prog[i]
				}
			}
		}
		for i := range argue {
			z01.PrintRune(rune(prog[i]) + "0")
			z01.PrintRune('\n')
		}
	}s
}
