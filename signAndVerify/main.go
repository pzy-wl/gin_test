package main

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	//content   := strings.Repeat("H", 244)+"e"
	//content   := strings.Repeat("H", 245)+"e"
	content := strings.Repeat("HBC", 10) + "e"
	//privateKey, publicKey := NewRsa("", "").CreateKeys(1024)
	fmt.Println("内容是：", content)
	privateKey, publicKey := NewRsa("", "").CreatePkcs8Keys(2048)
	fmt.Printf("公钥：%v\n私钥：%v\n", publicKey, privateKey)

	rsaObj := NewRsa(publicKey, privateKey)
	//使用接收方公钥对数据进行加密
	secretData, err := rsaObj.Encrypt([]byte("hello golang"))
	if err != nil {
		fmt.Println(err)
	}
	//对密文进行签名（使用自己的私钥）
	sign, _ := rsaObj.Sign(secretData, crypto.SHA1)
	fmt.Println("签名是：", sign)
	fmt.Println("密文是：", string(secretData))
	verify := rsaObj.Verify(secretData, sign, crypto.SHA1)
	//对机密数据进行解密
	if !verify {
		fmt.Println("验签失败 ，密文无效")
		return
	}
	fmt.Println("验签成功， 正在进解密")
	plainData, err := rsaObj.Decrypt(secretData)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf(" 加密：%v\n 解密：%v\n 签名：%v\n 验签结果：%v\n",
		hex.EncodeToString(secretData),
		string(plainData),
		hex.EncodeToString(sign),
		verify,
	)
}
