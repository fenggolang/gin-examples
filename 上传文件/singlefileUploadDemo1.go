package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 为multipart forms设置较低的内存限制(默认是32MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/upload", func(context *gin.Context) {
		// 单文件
		file, _ := context.FormFile("file")
		log.Println("上传的文件名称是:", file.Filename)

		// 上传文件至指定目录
		dst := "E:/" + file.Filename
		err := context.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Println("上传到服务器上的文件保存失败:", err.Error())
		}
		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.Run(":8080")
}
