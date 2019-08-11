package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32MiB)
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.Static("/", "./public")
	router.POST("/upload", func(context *gin.Context) {
		name := context.PostForm("name")
		email := context.PostForm("email")

		// Source
		file, err := context.FormFile("file")
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)
		if err := context.SaveUploadedFile(file, filename); err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		context.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fileds name=%s and email=%s", file.Filename, name, email))
	})
	router.Run(":8080")
}
