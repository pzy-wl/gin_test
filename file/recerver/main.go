package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件内容
func RecFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err =", err)
		return
	}

	buf := make([]byte, 1024*4)
	//读文件内容
	for {
		//注意这里,接收的是发送方发来的信息,不是上面读取的fileName的信息(f.Read)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("f.Read err =", err)
			}
			return
		}
		if n == 0 {
			fmt.Println("n == 0,文件接收完毕")
			break
		}
		f.Write(buf[:n]) //往文件写入内容
	}
}

func main() {
	//创建进行监听的套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	defer listener.Close()

	//创建用于通讯的套接字,阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err =", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf) //读取通讯方发送的文件名
	if err != nil {
		fmt.Println("conn.Read err =", err)
		return
	}

	fileName := string(buf[:n])
	//回复"ok"
	conn.Write([]byte("ok"))

	//接收文件内容
	RecFile(fileName, conn)
}
