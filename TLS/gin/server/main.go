package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
func main() {
	fmt.Println("GoLang 获取程序运行绝对路径")
	fmt.Println(GetCurrPath())
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello golang")
	})
	//命令行运行是此路径
	//router.RunTLS(":8080", "../../server.pem", "../../server.key")
	//点运行键执行此程序 文件路径是当前项目的根目录
	router.RunTLS(":8080", "./TLS/server.pem", "./TLS/server.key")
}
