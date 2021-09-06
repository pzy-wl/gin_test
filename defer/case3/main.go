package main

import "fmt"

//defer 调用的函数参数的值在 defer 定义时就确定了, 而 defer 函数内部所使用的变量的值需要在这个函数运行时才确定
func main() {
	i := 1
	defer fmt.Println("Deferred print:", i)
	i++
	fmt.Println("Normal print:", i)
	defer func(i int) {
		fmt.Println("deferred2 print:", i)
	}(i)
}
