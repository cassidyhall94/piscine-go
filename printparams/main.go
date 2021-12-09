package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argue := os.Args             // name a var to access the os.args command line
	for i, prog := range argue { // to loop over the command one by one
		if i != 0 { // if it does not start at the first index(0)
			for _, prog := range prog { // to loop over the command again after index (0)
				z01.PrintRune(prog) // print the command line without the first index (0)
			}
			z01.PrintRune('\n')
		}
	}
}
