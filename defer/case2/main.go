package main

import "fmt"

//defer的执行顺序：后进先出  程序发生panic,限制性defer后输出panic
func main() {
	deferCall()
}

func deferCall() {
	defer func() {
		fmt.Println("1111")
	}()

	defer func() {
		fmt.Println("2222")
	}()

	defer func() {
		fmt.Println("3333")
	}()

	panic("触发异常")
}
