package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"strings"
)

type info struct {
	id     string
	name   string
	birth  string
	depart string
}

func main() {
	files, err := ioutil.ReadDir("C:\\Users\\Administrator\\Desktop\\文件\\mail\\mile20210902")
	if err != nil {
		panic(err)
	}
	list := make([]*info, 0)
	// 获取文件，并输出它们的名字
	for _, file := range files {
		fmt.Println(file.Name())
		user := new(info)
		f, err := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\文件\\mail\\mile20210902\\" + file.Name())
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		//读取utf16编码格式的文件
		decoder := mahonia.NewDecoder("utf16")
		l := strings.Split(decoder.ConvertString(string(f)), "\n")
		for _, v := range l {
			if strings.Contains(v, "员工编码：") {
				index := strings.Index(v, "：")
				length := len("：")
				//fmt.Println("id是:  ", v[index+length:])
				user.id = v[index+length:]
				continue
			}
			if strings.Contains(v, "姓名：") {
				index := strings.Index(v, "：")
				length := len("：")
				//fmt.Println("姓名:  ", v[index+length:])
				user.name = v[index+length:]
				continue
			}
			if strings.Contains(v, "出生日期：") {
				index := strings.Index(v, "：")
				length := len("：")
				//fmt.Println("出生日期：", v[index+length:])
				user.birth = v[index+length:]
				continue
			}
			if strings.Contains(v, "部门：") {
				index := strings.Index(v, "：")
				length := len("：")
				//fmt.Println("部门： ", v[index+length:])
				user.depart = v[index+length:]
				continue
			}

		}
		list = append(list, user)
	}
	fmt.Println("正常用户数量有：", len(list))
	file := xlsx.NewFile()              // NewWriter 创建一个Excel写操作实例
	sheet, err := file.AddSheet("info") //表实例
	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, v := range list {
		row := sheet.AddRow()
		row.AddCell().SetValue(v.id)
		row.AddCell().SetValue(v.name)
		row.AddCell().SetValue(v.birth)
		row.AddCell().SetValue(v.depart)
	}
	file.Save("./info.xlsx")
}
