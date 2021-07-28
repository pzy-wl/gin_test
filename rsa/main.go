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
	content := strings.Repeat("H", 24270) + "e"
	//privateKey, publicKey := NewRsa("", "").CreateKeys(1024)
	privateKey, publicKey := NewRsa("", "").CreateKeys(2048)
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
