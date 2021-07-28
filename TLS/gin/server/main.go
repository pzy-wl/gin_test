package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello golang")
	})
	router.RunTLS(":8080", "../../ssl.pem", "../../ssl.key")
}
