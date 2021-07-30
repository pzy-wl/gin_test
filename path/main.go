package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Write([]byte("this is a test of golang path！"))
}
