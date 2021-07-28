// 一般用于网络协商数据加密的密钥
// 流程
// 1.生成本地的随机数
// 2.生成exchange key
// 3.交换exchange key
// 4.根据得到的exchange key，生成crypto key，作为aes加密
package ecdh

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var (
	//p = big.NewInt(115792089210356282266426440264866080066222482802048440882882484600088244664620)        // Diffie-Hellman交换密钥算法的质数
	//p = big.NewInt(4967644806815575240882406604862268688200240406680240828606208460266860824226420246862) // Diffie-Hellman交换密钥算法的质数
	p = big.NewInt(0)                               // Diffie-Hellman交换密钥算法的质数
	g = big.NewInt(0)                               // Diffie-Hellman交换密钥算法的质数
	r = rand.New(rand.NewSource(time.Now().Unix())) // 随机数
)

func init() {
	p.SetString("115792089210356248756420345214020892766250353991924191454421193933289684991999", 10) // 正好32个
	g.SetString("8179555366765431328676981059473238234208317050258169427898435860325837396910", 10)   // 正好32个
}

// 生成随机大数a，n表示生成的a存储大小
func RandomInt() *big.Int {
	var b [256]byte
	r.Read(b[:])
	a := big.NewInt(0).SetBytes(b[:])
	return a
}

// 生成交换密钥，a是本地的随机数，b是接收key的缓存
func GenExchangeKey(a *big.Int, b []byte) {
	fmt.Printf("p:%v g是：%v\n", p.String(), g.String())
	n := big.NewInt(0).Exp(g, a, p)
	n.FillBytes(b[:])
}

// 生成crypto key，exchangeKey是对方的交换密钥，cryptoKey是接收key的缓存，n是本地生成的随机数
func GenCryptoKey(exchangeKey, cryptoKey []byte, n *big.Int) {
	// 还原
	a := big.NewInt(0).SetBytes(exchangeKey)
	// 生成crypto key
	b := big.NewInt(0).Exp(a, n, p)
	b.FillBytes(cryptoKey[:])
}
