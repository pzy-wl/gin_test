package main

import "fmt"

//r的值先自增然后执行defer函数
func f1() (r int) {
	r = 1
	defer func() {
		r++
		fmt.Println(" r value = ", r)
	}()
	r++
	return
}

func main() {
	f1()
}
