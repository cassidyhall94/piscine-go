package main

import (
	"fmt"
	"os"
)

func main() {
	fname := "File name missing"
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("\n")
	} else {
		arr := make([]byte, 6)
		file.Read(arr)
		file.Close()
	}
}
