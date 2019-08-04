package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用LoadHTMLGlob()或者LoadHTMLFiles()
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html","templates/template2.html")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
