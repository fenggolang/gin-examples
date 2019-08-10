package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message") // Content-Type选择multipart/form-data，application/x-www-form-urlencoded; charset=UTF-8都可以
		nick := context.DefaultPostForm("nick", "anonymous")

		contentType := context.GetHeader("Content-Type")
		fmt.Println(contentType)
		context.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}
