package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 为multipart forms 设置较低的内存限制(默认是32MiB)
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/upload", func(context *gin.Context) {
		// Multipart form
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件至指定目录
			dst := "E:/" + file.Filename
			context.SaveUploadedFile(file, dst)
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":8080")
}
