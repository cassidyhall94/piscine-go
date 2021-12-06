package piscine

import "fmt"

func FirstRune(s string) rune {
	result := []rune(s)[0]
	for letter := range s {
		fmt.Printf("%c", letter)
	}
	return result
}
