package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//读取文件并且转化为字节数组
	path := "../数据模板.xlsx"
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("获取文件信息失败：", err.Error())
	}
	fmt.Println("文件名是：", info.Name())
	l, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("文件转化为字节码失败", err.Error())
		return
	}
	newfile1, err := os.Create("新文件1.xlsx")
	if err != nil {
		fmt.Println("新建文件失败", err.Error())
		return
	}
	_, err = newfile1.Write(l)
	if err != nil {
		fmt.Println("写入文件失败", err.Error())
		return
	}
}
