package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for _, check := range args {
		if check == "01" || check == "galaxy" || check == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
