package main

import (
	"os"
	"strconv"
)

func isSigns(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func main() {
	signs := []string{"+", "*", "-", "/", "%"}
	args := os.Args[1:]
	if len(args) != 3 {
	} else {
		if isSigns(args[1], signs) {
			nb1, err := strconv.Atoi(args[0])
			nb2, err2 := strconv.Atoi(args[2])
			if err == nil && err2 == nil {
				switch args[1] {
				case "+":
					os.Stderr.WriteString("nb1 + nb2")
				case "-":
					os.Stderr.WriteString("nb1 - nb2")
				case "/":
					if nb2 == 0 {
						os.Stderr.WriteString("No division by 0")
					} else {
						os.Stderr.WriteString("nb1 / nb2")
					}
				case "%":
					if nb2 == 0 {
						os.Stderr.WriteString("No Modulo by 0")
					} else {
						os.Stderr.WriteString("nb1 % nb2")
					}
				case "*":
					os.Stderr.WriteString("nb1 * nb2")

				}
			} else {
				os.Stderr.WriteString("1")
			}
		} else {
			os.Stderr.WriteString("0")
		}
	}
}
