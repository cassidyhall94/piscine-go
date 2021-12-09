package main

import (
	"os"

	"github.com/01-edu/z01"
)

func sliceruneSort() {
	// slicerune == num
	var slicerune [][]rune
	for _, arg := range os.Args[1:] { // to loop over the command one by one
		var rs []rune
		for _, argrune := range arg {
			rs = append(rs, argrune)
		}
		slicerune = append(slicerune, rs)
	}
	sortascii := len(slicerune)
	for i := 0; i < sortascii; i++ {
		for j := 0; j < (sortascii - 1 - i); j++ {
			if slicerune[j][0] > slicerune[j+1][0] {
				slicerune[j], slicerune[j+1] = slicerune[j+1], slicerune[j]
			}
		}
	}
	for _, sr := range slicerune {
		for _, j := range sr {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
