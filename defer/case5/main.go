package main

import "fmt"

//defer 函数调用的执行时机是外层函数设置返回值之后, 并且在即将返回之前
//return XXX 操作并不是原子的
func f3() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func main() {
	fmt.Println(f3())
}
