package model

import (
	"crypto/rand"
	"crypto/rsa"
)

//RsaComponent1 rsa公私钥对结构体
type RsaComponent1 struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

//Encrypt rsa加密
func (rc *RsaComponent1) Encrypt(plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, rc.PublicKey, plaintext)
}

//Decrypt rsa解密
func (rc *RsaComponent1) Decrypt(cipherText []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, rc.PrivateKey, cipherText)
}

//GenRsaKey 生成rsa私钥 参数bits: 指定生成的秘钥的长度, 单位: bit
func (rc *RsaComponent1) GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	publicKey := &privateKey.PublicKey
	rc.PublicKey = publicKey
	rc.PrivateKey = privateKey
	return err
}
