package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

func main() {
	r := gin.Default()

	// gin.H是map[string]interface{}额的一种快捷方式
	r.GET("/someJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	r.GET("moreJSON", func(context *gin.Context) {
		// 也可以使用一个结构体
		var msg struct {
			Name    string `json:"name"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		context.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})
	r.GET("/someYAML", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})
	r.GET("/someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf的具体定义写在testdata/protoexample文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被protoexample.Test protobuf序列化了的数据
		context.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}
