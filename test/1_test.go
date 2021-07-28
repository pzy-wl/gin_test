package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestZip(t *testing.T) {
	file, err := os.Create("test.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	zipwriter := zip.NewWriter(file)
	defer zipwriter.Close()

	fs := []struct {
		Url, Name string
	}{
		{Url: "http://videoactivity.bookan.com.cn/photo_10_202011301057385583397829_t.jpg", Name: "/1/11/1.jpg"},
		{Url: "http://videoactivity.bookan.com.cn/photo_46_202011292229184376672977_t.jpg", Name: "2.jpg"},
		{Url: "http://videoactivity.bookan.com.cn/photo_46_202011292229229158424627_t.jpg", Name: "3.jpg"},
		{Url: "http://videoactivity.bookan.com.cn/photo_46_202011292215023830368256_t.jpg", Name: "4.jpg"},
		{Url: "http://videoactivity.bookan.com.cn/photo_46_202011291755419090055749_t.jpg", Name: "5.jpg"},
		{Url: "http://videoactivity.bookan.com.cn/photo_46_202011291755420502426479_t.jpg", Name: "6.jpg"},
	}
	for _, f := range fs {
		iowriter, err := zipwriter.Create("")
		if err != nil {
			if os.IsPermission(err) {
				fmt.Println("权限不足: ", err)
				return
			}
			fmt.Printf("Create file %s error: %s\n", f.Name, err.Error())
			return
		}

		var content []byte
		resp, err := http.Get(f.Url)
		if err == nil {
			content, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				content = []byte("")
			}
			resp.Body.Close()
		}

		iowriter.Write(content)
	}

	// 创建空目录
	//zipwriter.Create("/name/dir/")
}

func TestKey(t *testing.T) {
	pub := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7cbdRD9npoEjXN8UFohO\nBTjKD4tYNvW73IaKXikcio6gnr9ulx2zZeAm0go7zKbbsJx7pZob623KrXsLw8UU\nz5nmAI+0Ww6d0PumWlqyQzOv0F1wtZ4MGROPvKvyjq0poLWtAsJVzaH5O52lKdOf\nUAuePDAzHp+WmB86zDWRAZWs25YdAQgmqBothGNKpX70BmT/T8qM5aPLG/PGo67g\nKNrk2DWzPbL8dnGLuwWZYL8QKv/ghQsA1X9krrrvyAEMGcJX/El6HbmhdiAZkam9\nNK6dLmbc2cRdB9/AJtUu+/mSJgni64/EsVqYGYhcTI/CoX8Wn0Ok4DZswASMDb8D\niwIDAQAB\n-----END PUBLIC KEY-----\n"
	pub1 := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4uVeMdIGsrcE5oLXqF2G\npEoyCdceLtbSoX0kN4vmHyJtQ89Z1azeR0vILPOaIjsHpNBTDmafVXEQjEkJCZH3\nQmil2XP3wGkGqCxOLUDKQ8HmDVX4fLcYrylYBwNL7zdF2xlAJEHoMIuu5nXxzY+J\nhT0Z6V/viwSIOxQHREtrdmIj7cGlaRsyQb60dcfev5ugP8988WUhJX/6c4a2Y2cV\nOJMtneSFvuxC/gBz6zvvT3EXT8/YUNrddCjLfb+vNukcewbgz/lh1XuCgyemQFiH\nkwKXE9bF2JQFd5HMw/9+FB1zfiORcSILQCfe4+DDHj6M1n2PbbkFS2+05f2tkeSP\nFwIDAQAB\n-----END PUBLIC KEY-----\n"
	fmt.Println("公钥对比结果为：", pub == pub1)
	pri := "5b551a04bce707bca73c97c7adc81cfc728b7ebb4baeaddbf89cd25bf0429169d21d71338107f92527462d4513c379996e06515e8dbb59bd32a81adbaeef9e59fd26d865680770ee44fc37d3d1c5f5b2552eafc9cc8a19a0968aefea1b4de438acc98774b10044cab9a50e61580ec20cacc8fcaf8d9c399d0050e97eed2f91741ab672901363aa5db4cec83ad0688acf27adb1cb43f3f7a435588520c3b8856e4afcb711911ed38a2463a552c89642b65f5c87ca383cdb043edeace4ddb4b10a2cc0d2a300c15d1ffd78f4c4870a98e9a22a182dff6f34973b6aa0808014c91d49b85b341d38cb138cd1228bbb5e27df71f20b979bec7e44c97c86218ee691cc"
	pri1 := "7b63d6ebff172784037ba499152961388aade142f7583e1e58c52a2e7931a40caf93109f3ebdefc1b64500b4167278f2a99e77f3323a1afc87ff3522ff6166017cb27b5da6663219c88d72a9ce0e01ae1cf80238738a4a87a9f1755b4949e6b837e6c5d855830847c7e9658c5ea9bbec7b0551c2b02c2a8d27a7ebaa0349f284581fa715b0feb38e613e3749957563526f99e140cd6c9c0149445be8fc39e849c5648835f23b3ca4b4fc47a5de788a6bbe69dc13a2c087ac2764ee47bed3c6843622d543f93d4719bcb9834599928f52c9dd5e10a087438ae9783e168cd378504d1a41f61158edbdd5e8706547d40747908470f3817942c6d0020149b5b9b86a"
	fmt.Println("aes对比结果为：", pri == pri1)
}

func TestLog(t *testing.T) {
	//	测试输出log常量 "|" 是按位或的意思
	fmt.Println(log.LstdFlags | log.Lshortfile)
	fmt.Println(log.LstdFlags + log.Lshortfile)
	fmt.Println(0 | 0)
}

func TestAtoi(t *testing.T) {
	i, err := strconv.Atoi("")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(i)
}

func TestGenG(t *testing.T) {
	//测试求一个大素数的本原根
	p, b := big.NewInt(1).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	if !b {
		fmt.Println("大素数赋值失败!")
	}
	println("开始")
	pBase := p
	p.Sub(p, big.NewInt(1))
	g := big.NewInt(1)
	for true {
		g.Add(g, big.NewInt(1))
		fmt.Printf("g:%v\n", g.String())
		gFlag := new(big.Int)
		gFlag.Exp(g, p, pBase)
		if 0 == gFlag.Cmp(big.NewInt(1)) {
			break
		}
	}
	fmt.Printf("p:%v\n", p.String())
	fmt.Printf("g:%v\n", g.String())
}

func TestAdd(t *testing.T) {
	a := big.NewInt(1)
	fmt.Printf("a:%v\n", a.String())
	for true {
		a.Add(a, big.NewInt(1))
		fmt.Printf("a:%v\n", a.String())
	}
}

func TestG(t *testing.T) {
	var pBaseHex string = "FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF"
	//var pBaseHex string = "FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AAAC42DAD33170D04507A33A85521ABDF1CBA64ECFB850458DBEF0A8AEA71575D060C7DB3970F85A6E1E4C7ABF5AE8CDB0933D71E8C94E04A25619DCEE3D2261AD2EE6BF12FFA06D98A0864D87602733EC86A64521F2B18177B200CBBE117577A615D6C770988C0BAD946E208E24FA074E5AB3143DB5BFCE0FD108E4B82D120A93AD2CAFFFFFFFFFFFFFFFF"
	p := new(big.Int)
	p, _ = p.SetString(pBaseHex, 16)
	pBase := new(big.Int)
	pBase, _ = pBase.SetString(pBaseHex, 16)
	p.Sub(p, big.NewInt(1))
	// 计算g
	g := new(big.Int)
	for {
		g.Rand(rand.New(rand.NewSource(time.Now().Unix())), p)
		gFlag := new(big.Int)
		gFlag.Exp(g, p, pBase)
		if 0 == gFlag.Cmp(big.NewInt(1)) {
			break
		}
	}
	// 随机出server_number
	serverNumber := new(big.Int)
	serverRandMax := big.NewInt(10000000)
	serverNumber.Rand(rand.New(rand.NewSource(time.Now().Unix())), serverRandMax)
	// 最终的processed_server_number
	processedServerNumber := new(big.Int)
	processedServerNumber = processedServerNumber.Exp(g, serverNumber, pBase)
	// 返回给客户端数据
	fmt.Printf("g is :%v\n", g.String())
	fmt.Printf("g length is :%v\n", len(g.String()))
	fmt.Printf("p is :%v\n", pBase.String())
	fmt.Printf("p length is :%v\n", len(p.String()))
}
