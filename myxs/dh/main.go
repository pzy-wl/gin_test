package ecdh

import (
	"bytes"
	"fmt"
	"math/big"
)

func main() {
	privateKey1, publicKey1 := NewRsa("", "").CreatePkcs8Keys(2048)
	privateKey2, publicKey2 := NewRsa("", "").CreatePkcs8Keys(2048)
	fmt.Printf("公钥：%v\n私钥：%v\n", publicKey1, privateKey1)
	a1 := big.NewInt(0).SetBytes([]byte(privateKey1))
	a2 := big.NewInt(0).SetBytes([]byte(privateKey2))
	fmt.Printf("公钥：%v\n私钥：%v\n", a1.Int64(), a2.Int64())
	key1 := make([]byte, 2048)
	GenCryptoKey([]byte(publicKey2), key1, a1)
	key2 := make([]byte, 20480)
	GenCryptoKey([]byte(publicKey1), key2, a2)
	if bytes.Compare(key1, key2) != 0 {
		fmt.Println("协商密钥失败!")
		fmt.Println("key1:", big.NewInt(0).SetBytes(key2))
		fmt.Println("key1:", big.NewInt(0).SetBytes(key1))
		return
	}
	fmt.Printf("双方生成的密钥同， 都是:%v\n", big.NewInt(0).SetBytes(key1).String())
}
