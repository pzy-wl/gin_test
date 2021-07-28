package ecdh

import (
	"bytes"
	"crypto"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
)

func TestRsa(t *testing.T) {
	content := strings.Repeat("H", 24270) + "e"
	//privateKey, publicKey := NewRsa("", "").CreateKeys(1024)
	privateKey, publicKey := NewRsa("", "").CreatePkcs8Keys(2048)
	fmt.Printf("公钥：%v\n私钥：%v\n", publicKey, privateKey)

	rsaObj := NewRsa(publicKey, privateKey)
	secretData, err := rsaObj.Encrypt([]byte(content))
	if err != nil {
		fmt.Println(err)
	}
	plainData, err := rsaObj.Decrypt(secretData)
	if err != nil {
		fmt.Print(err)
	}

	data := []byte(strings.Repeat(content, 200))
	//sign,_ := rsaObj.Sign(data, crypto.SHA1)
	//verify := rsaObj.Verify(data, sign, crypto.SHA1)

	sign, _ := rsaObj.Sign(data, crypto.SHA256)
	verify := rsaObj.Verify(data, sign, crypto.SHA256)

	fmt.Printf(" 加密：%v\n 解密：%v\n 签名：%v\n 验签结果：%v\n",
		hex.EncodeToString(secretData),
		string(plainData),
		hex.EncodeToString(sign),
		verify,
	)
}

func TestRsaDh(t *testing.T) {
	privateKey1, publicKey1 := NewRsa("", "").CreatePkcs8Keys(2048)
	privateKey2, publicKey2 := NewRsa("", "").CreatePkcs8Keys(2048)
	fmt.Printf("公钥：%v\n私钥：%v\n", publicKey1, privateKey1)
	a1 := big.NewInt(0).SetBytes([]byte(privateKey1))
	a2 := big.NewInt(0).SetBytes([]byte(privateKey2))
	fmt.Printf("公钥：%v\n私钥：%v\n", a1.Int64(), a2.Int64())
	key1 := make([]byte, 2048)
	GenCryptoKey([]byte(publicKey2), key1, a1)
	key2 := make([]byte, 2048)
	GenCryptoKey([]byte(publicKey1), key2, a2)
	if bytes.Compare(key1, key2) != 0 {
		fmt.Println(key1)
		fmt.Println(key2)
	}
	fmt.Printf("双方生成的密钥同， 都是:%v\n", big.NewInt(0).SetBytes(key1).String())
}
