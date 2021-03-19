package main

import (
	_ "ginTest/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @Summary 打印测试功能
// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1
// @Host 127.0.0.1:8080
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {string} json "{"code":200,"data":"name","msg":"ok"}"
// @Router / [get]
func Print(context *gin.Context) {
	var (
		name string
	)
	name = context.Query("name")
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": name,
	})
}

func main() {
	var (
		route *gin.Engine
		v1    *gin.RouterGroup
	)
	route = gin.Default()
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//gin.SetMode(gin.ReleaseMode)
	v1 = route.Group("/api/v1")
	{
		v1.GET("/", Print)
	}

	route.Run()
}
