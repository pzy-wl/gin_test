package ecdh

import (
	"bytes"
	"crypto"
	"fmt"
	"math/big"
	"testing"
)

//双方及逆行密钥协商
func Test_DiffieHellman(t *testing.T) {
	keyLen := 256
	// 1. 生成本地的随机数
	clientRandom := RandomInt()
	serverRandom := RandomInt()
	fmt.Println("clientRandom:", clientRandom)
	// 2. 生成交换密钥
	clientExchangeKey := make([]byte, keyLen)
	GenExchangeKey(clientRandom, clientExchangeKey)
	serverExchangeKey := make([]byte, keyLen)
	GenExchangeKey(serverRandom, serverExchangeKey)
	// 3. 生成加密密钥
	clientCryptoKey := make([]byte, keyLen)
	GenCryptoKey(serverExchangeKey, clientCryptoKey, clientRandom)
	serverCryptoKey := make([]byte, keyLen)
	GenCryptoKey(clientExchangeKey, serverCryptoKey, serverRandom)
	// 比较双方生成的密钥，应该是一样的
	if bytes.Compare(clientCryptoKey, serverCryptoKey) != 0 {
		t.FailNow()
	}
	fmt.Printf("客户端的交换秘钥是：%v\n", big.NewInt(0).SetBytes(clientExchangeKey).String())
	fmt.Printf("服务端的交换秘钥是：%v\n", big.NewInt(0).SetBytes(serverExchangeKey).String())
	fmt.Printf("双方生成的密钥同， 都是:%v\n", big.NewInt(0).SetBytes(clientCryptoKey).String())
}

//多方方及逆行密钥协商
func Test_DiffieHellman1(t *testing.T) {
	keyLen := 32
	// 1. 生成本地的随机数
	Random1 := RandomInt()
	fmt.Printf("%v\n", Random1)
	Random2 := RandomInt()
	fmt.Printf("%v\n", Random2)
	Random3 := RandomInt()
	fmt.Printf("%v\n", Random3)
	// 2. 生成交换密钥
	client1 := make([]byte, keyLen)
	GenExchangeKey(Random1, client1)
	client2 := make([]byte, keyLen)
	GenExchangeKey(Random2, client2)
	client3 := make([]byte, keyLen)
	GenExchangeKey(Random3, client3)
	// 3. 生成加密密钥
	//client1生成
	client12 := make([]byte, keyLen)
	GenCryptoKey(client1, client12, Random2)
	client123 := make([]byte, keyLen)
	GenCryptoKey(client12, client123, Random3)

	//client2密钥生成
	client23 := make([]byte, keyLen)
	GenCryptoKey(client2, client23, Random3)
	client231 := make([]byte, keyLen)
	GenCryptoKey(client23, client231, Random1)

	//client3密钥生成
	client31 := make([]byte, keyLen)
	GenCryptoKey(client3, client31, Random1)
	client312 := make([]byte, keyLen)
	GenCryptoKey(client31, client312, Random2)

	// 比较三方生成的密钥，应该是一样的
	if bytes.Compare(client231, client123) != 0 {
		t.FailNow()
	}
	if bytes.Compare(client231, client312) != 0 {
		t.FailNow()
	}
	if bytes.Compare(client123, client312) != 0 {
		t.FailNow()
	}
	fmt.Println("三方生成的密钥同")
	//
	//// 有一个黑客
	//hackerRandom := RandomInt(keyLen)
	//// 获取到了server的交换key
	//hackerCryptoKey := make([]byte, keyLen)
	//GenCryptoKey(server2, hackerCryptoKey, hackerRandom)
	//// 要冒充，生成的密钥是不一致的
	//if bytes.Compare(server2, hackerCryptoKey) == 0 {
	//	t.FailNow()
	//}
}

func TestExp(t *testing.T) {
	//测试模幂运算
	a := big.NewInt(2)
	b := big.NewInt(2)
	c := big.NewInt(5)
	fmt.Println("模幂运算结果是：", big.NewInt(0).Exp(a, b, c))
}

func TestBig(t *testing.T) {
	a := big.NewInt(0).SetBytes([]byte("-----BEGIN PRIVATE KEY-----\nMIIEvAIBADALBgkqhkiG9w0BAQEEggSoMIIEpAIBAAKCAQEA7eof/0upBN5tbmxg\nJdDiLr88R+fdY8hP4wLOFNYcoHAOse9vqKUYKtyU76QVdNbqO2w4gwCvzEezVbJK\n3qccukz1gp0gG2SsYub8WvH+vpUS5hXWJsd4WE443VldeV1B+hKHUh47ar6PhDEq\n4YgXgwyijZd1NIKJ+xPtxAoF0VOul0GF8RuweaTvEN4C33HTyaV01yb2cXbTo96b\nWA+8FKLNOR7gmGEHCdIyX8BD+z1oaEPFG3DMityfHdDdDAth0m4mFCaVEtBMOBDJ\n/uNPnIyFxp612jsH+HlnhaXVklc7BkrPbDTO+IyvOf4R9pYVjXJ0mKYw8V5NVOBZ\nnWGlowIDAQABAoIBAFVQaIr2ybRKQrmfOVc2MXWL9ATg+33FnggMwHOuV/OcW1ip\nQMQb96+fC/VnRJ5yVupaI2WuwFujtoZbegefq2iPIlblG1dXYY3RwNqn/q6+7Fj/\nqZycWOZpnnCZlDBqJbeUH1xuJOZhUTuZGgRn9e5zgHL/xfK/gNU9TzzGT1HBXEpF\nin6u10sz7bKgFM/wfXF9C92ZiMkqWVRe7R3gi80HUyu79+sHUXTFsnbeG3y+e3NZ\nZW8mCafauS1VjgPUIJK7OAEKYZeNyymQIfk63WBSVfFmIHyVLgSykLG9EZIGbebz\nO1krrU5DZO1VnWnBezQNWjcwXvYVYY3lRcDDwekCgYEA+1sijU8BsOELLhC/Bk16\nZHnxoTfjlEsVwKffOalJN+1q+aQY2sQeuH+pHs4TBfNzgb/SumK33JiHiQuMVLvI\nyrVWSMENycqdHY7Ov/8dDXXfTGmbby2tgWBLbo8XVRe5RfgEprdHRQD3ruNRg3hb\nLbIpChgKbiA/ZrzexWIlYx8CgYEA8k9qIvNfKMhakFWwbV5Z+yLWz8URqbBDZVQp\n+rkHktpKDG7PD0jQLIfAW6SUvV99YqgVcZcwKjyvVe0oW+XgQ6mpJkBeCPLpyYwK\nEG7ai7XkX4v2JqxRGxSSsn3zQprWxtfGLtiaScvT14OUo1Pg8IhZ5PqmITCxiZXD\nmc7rUP0CgYADc6jxz9fyGeURaOF2sLGxl4sd76qMasuNrcB/BjgfMH97lilRN8hE\nOFyWM7ZHKqS3b27xaFGmEeIqm6H9t+Z2Ai+BKEWMM+Ace82TDWnUuX2rive7eyEa\n2buvdP+ZB45fy5wkvfeYlv7PxAdmisyVwBCYNuRNxtuYCZf1lfd1eQKBgQC1GP91\nUbl8TInLOA/dHdhYZwhvTpVN4ovsboPVfHaZngYb1rFwdfYIYgk1NQNpqbWXh5JH\nTDzWRHay3MX+MQd06peeHBtryslKNhzLfV5fbrlZY8y7yKvxdmXRe2rVC86b3Gal\nWoOtxeX1O8VacHU8sFunVGHXioIWF1WSXeDnrQKBgQDiweVY7cwh1HjrAtxUAPo8\nIDLLiHjToeu1PdByfsPrB5zSEdCgxiigvtXSNjvEvmfYdKDqMrR6crHAN+P5SxoK\nANAA7NIKDTYluMsy4u1YQ2gHPJ42iSNG/+/3L5khS2WvaZCjzxy72DgajdljqIO5\ndBANZUCmvv10eiGgWewGnw==\n-----END PRIVATE KEY-----"))
	fmt.Printf("goEncrypt is :%v", a.Int64())
}

func TestEqual(t *testing.T) {
	fmt.Println("106697219132480173106064317148705638676529121742557567770857687729397446898790451577487723991083173"+
		"0102424168632380997160447756586819818214079227220527789589428918310335124632627410539616815129082180038404085269"+
		"15629689432111480588966800949428079015682624591636010678691927285321708935076221951173426894836169" == "106697219"+
		"13248017310606431714870563867652912174255756777085768772939744689879045157748772399108317301024241686323809971604477"+
		"565868198182140792272205277895894289183103351246326274105396168151290821800384040852691562968943211148058896680094942807"+
		"9015682624591636010678691927285321708935076221951173426894836169")

}

func TestSignAndVerity(t *testing.T) {
	//DH验证签名与验签
	keyLen := 32
	//生成私钥
	Random1 := RandomInt()
	client1 := make([]byte, keyLen)
	//生成公钥
	GenExchangeKey(Random1, client1)
	rsa := NewRsa(Random1.String(), big.NewInt(0).SetBytes(client1).String())
	data := []byte("hello, golang")
	sign, err := rsa.Sign(data, crypto.SHA1)
	if err != nil {
		fmt.Printf("sign err%v", err.Error())
		return
	}
	verity := rsa.Verify(data, sign, crypto.SHA1)
	fmt.Println("验证结果是：", verity)
}

func TestEnAndDe(t *testing.T) {
	//测试DH算法的加密与解密

}
