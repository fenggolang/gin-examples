package main

import (
	"github.com/gin-gonic/gin"
)

// Multipart/Urlencoded绑定

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

/**
测试：
curl -v --form user=wpc --form password=wpc@123 http://127.0.0.1:8080/login
post请求的时候，既可以使用Query,也可以使用Body,且选择对应的form-data和x-www-form-urlencoded程序都可以处理
*/
func main() {
	router := gin.Default()
	router.POST("/login", func(context *gin.Context) {
		// 你可以使用显示绑定声明绑定 multipart form:
		//var form LoginForm
		//context.ShouldBindWith(&form,binding.Form)

		// 或者简单地使用ShouldBind 方法自动绑定:
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if context.ShouldBind(&form) == nil {
			if form.User == "wpc" && form.Password == "wpc@123" {
				context.JSON(200, gin.H{
					"status": "登录成功"})
			} else {
				context.JSON(401, gin.H{
					"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080")
}
