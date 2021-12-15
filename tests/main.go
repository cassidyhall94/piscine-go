package main

import (
	"fmt"
	"piscine"
)

func main() {
	steps := piscine.CollatzCountdown(16)
	fmt.Println(steps)
}
