package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	G = 5
	N = 23
)

func main() {

	go clientA(0, 0)
	go clientB(0, 0)
	go clientC(0, 0)
	select {}
}

var a, b, c int

func clientA(num float64, step int64) {
	if step == 0 {
		a1 := math.Mod(math.Pow(G, float64(a)), N)
		clientB(a1, 1)
	}
	if step == 1 {
		//计算p^b^a%N
		a2 := math.Mod(math.Pow(num, float64(a)), N)
		clientB(a2, 2)
	}
	if step == 2 {
		//最终计算所得出的秘钥
		a3 := math.Mod(math.Pow(num, float64(a)), N)
		fmt.Println("a 计算得出的秘钥是", a3)
	}
	select {}
}

func clientB(num float64, step int64) {
	if step == 0 {
		b1 := math.Mod(math.Pow(G, float64(b)), N)
		clientC(b1, 1)
	}

	if step == 1 {
		//计算p^b^a%N
		b2 := math.Mod(math.Pow(num, float64(b)), N)
		clientC(b2, 2)
	}
	if step == 2 {
		//最终计算所得出的秘钥
		b3 := int64(math.Mod(math.Pow(num, float64(b)), N))
		fmt.Println("b 计算得出的秘钥是", b3)
	}
	select {}
}

func clientC(num float64, step int64) {
	if step == 0 {
		c1 := math.Mod(math.Pow(G, float64(c)), N)
		clientA(c1, 1)
	}
	if step == 1 {
		//计算p^b^a%N
		c2 := math.Mod(math.Pow(num, float64(c)), N)
		clientA(c2, 2)
	}
	if step == 2 {
		//最终计算所得出的秘钥
		c3 := int64(math.Mod(math.Pow(num, float64(c)), N))
		fmt.Println("c 计算得出的秘钥是", c3)
	}
	select {}
}

func init() {
	fmt.Println("正在执行初始化。。。。。")
	//初始化a,b, c
	//__________________a___________________
	rand.Seed(time.Now().Unix())
	a = rand.Intn(100)
	fmt.Println(a, "A")

	//-----------------B--------------
	time.Sleep(time.Second)
	rand.Seed(time.Now().Unix())
	b = rand.Intn(100)
	fmt.Println(b, "B")

	//________________c___________
	time.Sleep(time.Second)
	rand.Seed(time.Now().Unix())
	c = rand.Intn(100)
	fmt.Println(c, "C")
}
