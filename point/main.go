package main

import "fmt"

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21

	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
