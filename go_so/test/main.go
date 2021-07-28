package main

import (
	"fmt"
	"plugin"
)

func main() {
	ptr, err := plugin.Open("aplugin.so")
	if err != nil {
		fmt.Println("读取文件错误")
	}
	add, err := ptr.Lookup("Add")
	if err != nil {
		fmt.Println("读取函数错误")
	}
	sum := add.(func(int, int) int)(1, 2)
	fmt.Println("结果是：", sum)
}
