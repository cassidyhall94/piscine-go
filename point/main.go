package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	values := "x = 42, y = 21"

	for _, printpoint := range values {
		z01.PrintRune(printpoint)
	}
	z01.PrintRune('\n')
}
