package main

import (
	"fmt"
	"piscine"
)

func main() {
	steps := piscine.CollatzCountdown(5)
	fmt.Println(steps)
}
