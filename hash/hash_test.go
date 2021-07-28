package hash

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestHash1(t *testing.T) {
	//测试字符串哈希哈希
	message := []byte("hello world")
	hashCode := GetSHA256HashCode(message)
	fmt.Println(hashCode)
}

func TestHash2(t *testing.T) {
	//测试文件哈希
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("打开文件出错%v\n", err.Error())
		return
	}
	defer file.Close()
	body, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("读取文件出错%v\n", err.Error())
		return
	}
	fmt.Printf("得到的哈希值是%v\n", GetSHA256HashCode(body))
	err = os.Rename("./file/t.txt", "./file/test.txt")
	if err != nil {
		fmt.Printf("文件重命名失败%v", err.Error())
		return
	}
	b, err := ioutil.ReadFile("./file/test.txt")
	if err != nil {
		fmt.Printf("读取文件出错%v", err.Error())
		return
	}
	fmt.Printf("得到的哈希值是%v\n", GetSHA256HashCode(b))
}

func TestHash3(t *testing.T) {
	b, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Printf("读取文件出错%v", err.Error())
		return
	}
	hash := GetSHA256HashCode(b)
	fmt.Printf("得到的哈希值是%v\n", hash)
	fmt.Println(hash == "c2cc9cd148736a31e959b418fe7a1193c585a0189c72599ff6f5af57be325994")
}
