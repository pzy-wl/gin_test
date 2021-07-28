package main

import (
	"fmt"
	"github.com/wumansgy/goEncrypt"
	"math"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestNum(t *testing.T) {
	//	计算p^a^b^c%N
	//fmt.Println("2^3:", math.Pow(2, 3))
	a1 := math.Mod(math.Pow(5, 4), 23)
	b1 := math.Mod(math.Pow(5, 3), 23)
	fmt.Println("A1结果是：", a1)
	fmt.Println("B1结果是：", b1)
	b2 := math.Mod(math.Pow(a1, 3), 23)
	a2 := math.Mod(math.Pow(b1, 4), 23)
	fmt.Println("A2是：", a2)
	fmt.Println("B2是：", b2)

}

func TestEcc(t *testing.T) {
	goEncrypt.GetEccKey()
}

func TestGenG(t *testing.T) {
	//测试求一个大素数的本原根
	p, b := big.NewInt(1).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	if !b {
		fmt.Println("大素数赋值失败!")
	}
	pBase := new(big.Int)
	g := new(big.Int)
	for {
		g.Rand(rand.New(rand.NewSource(time.Now().Unix())), p)
		gFlag := new(big.Int)
		gFlag.Exp(g, p, pBase)
		if 0 == gFlag.Cmp(big.NewInt(1)) {
			break
		}
	}
	fmt.Printf("p:%v\n", p.String())
	fmt.Printf("g:%v\n", g.String())
}
