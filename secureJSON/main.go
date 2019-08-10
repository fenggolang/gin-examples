package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。
func main() {
	r := gin.Default()
	// 你也可以使用自己的SecureJSON前缀
	//r.SecureJsonPrefix(")]},\n")

	r.GET("/someJSON", func(context *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出while(1);["lena","austin","foo"]
		context.SecureJSON(http.StatusOK, names)
	})

	// 监听并在0.0.0.0:8080上启动服务
	r.Run(":8080")
}
