package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	// "fmt"
)

func main() {
	httpsRouter := gin.Default()
	httpRouter := gin.Default()

	httpsRouter.Use(TlsHandler()) // 处理SSL的中间
	//httpsRouter.StaticFS("/", http.Dir("./static"))
	//httpRouter.StaticFS("/", http.Dir("./static"))
	// 这部分根据实际业务逻辑进行修改
	go httpsRouter.RunTLS(":9999", "./ssl/ssl.pem", "./ssl/ssl.key")
	// 上面一行三个参数分别是SSL的监听端口，证书以及私钥
	httpRouter.Run(":9998")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
