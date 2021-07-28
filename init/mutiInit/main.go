package main
import "fmt"
import _ "net/http/pprof" //导入包但不使用，只是用其init函数`_`
//init函数比较特殊，可以在包里被多次定义。

//init函数的主要用途：初始化不能使用初始化表达式初始化的变量
func init() {
	fmt.Println("init 1")
}
func init() {
	fmt.Println("init 2")
}

var initArg [20]int
func init() {
	initArg[0] = 10
	for i := 1; i < len(initArg); i++ {
		initArg[i] = initArg[i-1] * 2
	}
}
func main() {
	fmt.Println("main")
}