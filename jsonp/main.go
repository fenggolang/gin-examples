package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP?callback=x", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// callback 是x
		// 将输出:
		context.JSONP(http.StatusOK, data)
	})

	// 监听并在0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
