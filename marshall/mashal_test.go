package marshall

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	//测试测试unmarshal
	name := "hello world"
	n, err := json.Marshal(name)
	if err != nil {
		fmt.Println("this is an error")
		return
	}
	l := make([]string, 0)
	json.Unmarshal(n, l)
	fmt.Println("切片是：", l)
	for _, v := range l {
		fmt.Println(v)
	}
}

func TestMarshall(t *testing.T) {
	//反序列化
	l := make([]string, 0)
	b, _ := json.Marshal("admin")
	err := json.Unmarshal(b, &l)
	fmt.Printf("err is %v", err.Error())
	fmt.Printf("err is %v", l)
}

func TestStringSlice(t *testing.T) {
	//切片转成string 格式并且unmarshall
	l := make([]string, 0)
	l = append(l, "admin")
	b, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err.Error())
	}
	l1 := make([]string, 0)
	fmt.Println(b)
	json.Unmarshal([]byte(string(b)), &l1)
	for i, s := range l1 {
		fmt.Println("正在遍历数组", i, s)
	}
	fmt.Printf("%v", &l1)
}
