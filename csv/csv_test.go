package main

import (
	"encoding/csv"
	"fmt"
	"ginTest/utils"
	"io"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestCsvReade(t *testing.T) {
	//	测试读取csv文件首行
	f, err := os.Open("./newfile.csv")
	if err != nil {
		fmt.Printf("读取文件时出错%v", err.Error())
		return
	}
	r := csv.NewReader(f)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Record has %d columns.\n", len(record))
		//city, _ := iconv.ConvertString(record[2], "gb2312", "utf-8")
		//
		//fmt.Printf("%s %s %s \n", record[0], record[1], city)
		for _, v := range record {
			fmt.Println(v)
		}
	}
}

//测试模板验证
func TestCompareDate(t *testing.T) {

	err := utils.UnZip("newfile1.zip", "", "./file")
	if err != nil {
		fmt.Printf("文件解压失败2！%v", err)
		return
	}
	f1, err := os.Open("./file/newfile.csv")
	if err != nil {
		fmt.Printf("打开csv文件失败! %v", err)
		return
	}
	r1 := csv.NewReader(f1)
	record2, err := r1.Read()
	if err != nil {
		fmt.Printf("文件读取失败！%v", err)
		return
	}
	err = utils.UnZip("newfile.zip", "", "./file1")
	if err != nil {
		fmt.Printf("文件解压失败1！%v", err.Error())
		return
	}
	f, err := os.Open("./file1/newfile.csv")
	if err != nil {
		fmt.Printf("打开csv文件失败! %v", err.Error())
		return
	}
	r := csv.NewReader(f)
	// Read each record from csv
	record1, err := r.Read()
	if err != nil {
		fmt.Printf("文件读取失败！%v", err.Error())
		return
	}
	fmt.Printf("数据模板的比较结果是%v", record1)
	fmt.Printf("数据模板的比较结果是%v\n", record2)
	println("比较结果是：", reflect.DeepEqual(record1, record2))
}

//测试数组比较
func TestSliceCompare(t *testing.T) {
	//支持比较，只支持 == 或 !=, 比较是不是每一个元素都一样，2个数组比较，数组类型要一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	fmt.Println(" a == b ", a == b)
	fmt.Println(" a == c ", a == c)
}
