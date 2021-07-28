package main

import "fmt"

func GetMoney() {
	fmt.Println("money")
	return
}

func main() {
	//GetMoney()
	i := 100
	fmt.Printf("I 的地址是：%v\n", &i)
}
