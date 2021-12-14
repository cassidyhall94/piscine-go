package main

import (
	"fmt"
	"os"
)

func main() {
	filearg := os.Args
	for _, a := range filearg {
		file, err := os.Open(a)
		if err != nil {
			fmt.Printf("the mistake is : %v\n", err.Error())
		} else {
			fmt.Println(file.Stat())
			arr := make([]byte, 6)
			file.Read(arr)
			fmt.Println(string(arr))
			file.Close()
		}
	}
}
