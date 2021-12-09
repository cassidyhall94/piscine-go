package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// slicerune == num
	var slicerune []rune
	for _, arg := range os.Args[1:] { // to loop over the command one by one
		r := rune(arg[0])
		slicerune = append(slicerune, r)
	}
	sortascii := len(slicerune)
	for i := 0; i < sortascii; i++ {
		for j := 0; j < (sortascii - 1 - i); j++ {
			if slicerune[j] > slicerune[j+1] {
				slicerune[j], slicerune[j+1] = slicerune[j+1], slicerune[j]
			}
		}
	}
	for i := range slicerune {
		z01.PrintRune(rune(slicerune[i]))
		z01.PrintRune('\n')
	}
}
