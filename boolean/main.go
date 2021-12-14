package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arg := os.Args
	for _, s := range arg {
		for _, r := range s {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := 0
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
