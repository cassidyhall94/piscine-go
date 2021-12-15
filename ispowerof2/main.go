package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		print("false\n")
		return
	}

	if n == 0 {
		print("false\n")
		return
	}

	if isPower(n) == 0 {
		print("true\n")
	} else {
		print("false\n")
	}
}

func isPower(n int) int {
	return n & (n-1)
}

func print(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
