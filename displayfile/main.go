package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("Too many arguments\n")
		return
	} else if len(os.Args) < 2 {
		fmt.Printf("File name missing\n")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("File name missing\n")
	}
	arr := make([]byte, 20)
	file.Read(arr)
	fmt.Println(string(arr))
	file.Close()
}
