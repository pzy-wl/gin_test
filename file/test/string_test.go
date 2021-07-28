package main

import (
	"fmt"
	"testing"
)

func TestStringToByte(t *testing.T) {
	//字符串转化为字节数组就是把每一个字符转换为其对应的ASCII码,一个汉字占三个字节， 一个中文标点占三个字节
	s := "汉字占三个字节，中文标点占也是占三个字节"
	fmt.Printf("原字符串是%s, 长度是%d\n", s, len(s))
	byte := []byte(s)
	fmt.Printf("转换后的字节数组是%v, 长度是%d\n", byte, len(byte))
}
