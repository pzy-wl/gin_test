package main

import (
	"encoding/json"
	"fmt"
)

type Persion struct {
	Name   string
	Age    int
	Gender string
	Job    string
}

func main() {
	tmp := new(Persion)
	tmp.Age = 12
	tmp.Name = "zs"
	tmp.Gender = "男"
	tmp.Job = "程序员"
	a1, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println("序列化时出错", err.Error())
	}
	fmt.Printf("序列化结果是：%v， 长度是%d\n", a1, len(a1))
	a2 := GZipBytes(a1)
	fmt.Printf("压缩后的结果是：%v, 长度是%d\n", a2, len(a2))
	a3 := UGZipBytes(a2)
	fmt.Printf("解压后的结果是：%v, 长度是%d\n", a3, len(a3))
	p := &Persion{}
	err = json.Unmarshal(a3, p)
	if err != nil {
		fmt.Println("反序列化时出错：", err.Error())
	}
	fmt.Printf("序列化结果是：%v", *tmp)
}
