package main

import (
	"flag"
	"fmt"
	"plugin"
)

var a = flag.Int("a", 0, "add element a+b:a")
var b = flag.Int("b", 0, "add element a+b:b")
var c = flag.Int("c", 0, "sub element c-d:c")
var d = flag.Int("d", 0, "sub element c-d:d")

func main() {
	flag.Parse()
	//fmt.Println("a:", *a)
	//fmt.Println("b:", *b)
	//fmt.Println("c:", *c)
	//fmt.Println("d:", *d)
	ptr, err := plugin.Open("aplugin.so")
	if err != nil {
		fmt.Println(err)
	}

	Add, _ := ptr.Lookup("Add")
	sum := Add.(func(int, int) int)(*a, *b)
	fmt.Println("Add结果：", sum)

	Sub,_ := ptr.Lookup("Subtract")
	sub := Sub.(func(int,int)int)(*c,*d)
	fmt.Println("Sub结果：",sub)
}