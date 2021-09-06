package main

//defer定于延迟调用，无论函数是否出错，它都确保结束前被调用
func expect(a, b int) {
	defer println("defering ...........")
	println(a / b)
}

func main() {
	expect(10, 0)
}
