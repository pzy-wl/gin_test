package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	_, error := os.Open("./main")
	if error != nil {
		fmt.Sprintf("123")
		return
	}
	fmt.Println("读取文件成功")
}
