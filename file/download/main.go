package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	engin := gin.Default()
	engin.GET("/downloadFiles", DownloadFileService)
	engin.GET("/", HELLO)
	engin.Run(":8080")
}

func HELLO(c *gin.Context) {
	c.JSON(200, "ok")
}

//TODO Test资源文件下载
func DownloadFileService(c *gin.Context) {
	fileDir := c.Query("fileDir")
	fileName := c.Query("fileName")
	fmt.Println("文件位置是：", fileDir+fileName)
	//打开文件
	_, errByOpenFile := os.Open(fileDir + fileName)
	//非空处理
	if errByOpenFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "失败",
			"error":   "资源不存在",
		})
		//c.Redirect(http.StatusFound, "/404")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(fileDir + fileName)
	return
}
