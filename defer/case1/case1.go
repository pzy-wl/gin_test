package main

import (
	"fmt"
)

//代码的延迟顺序与最终的执行顺序是反向的。
//延迟调用是在 defer 所在函数结束时进行，函数结束可以是正常返回时，也可以是发生宕机时。
func main() {
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")
}
