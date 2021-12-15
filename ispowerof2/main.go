package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 2 {
		return
	} else {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			for _, r := range "false\n" {
				z01.PrintRune(r)
			}
			return
		}

		if n == 0 {
			for _, r := range "false\n" {
				z01.PrintRune(r)
			}
			return
		}

		if isPower(n) == 0 {
			for _, r := range "true\n" {
				z01.PrintRune(r)
			}
		} else {
			for _, r := range "false\n" {
				z01.PrintRune(r)
			}
		}
	}
}

func isPower(n int) int {
	return n & (n - 1)
}
