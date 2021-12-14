package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Printf("File name missing\n")
	} else {
		arr := make([]byte, 20)
		file.Read(arr)
		fmt.Println(string(arr))
		file.Close()
	}
}