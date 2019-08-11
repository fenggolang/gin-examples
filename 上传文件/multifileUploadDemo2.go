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

		// Multipart form
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := context.SaveUploadedFile(file, filename); err != nil {
				context.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		context.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fileds name=%s and email=%s", len(files), name, email))
	})
	router.Run(":8080")
}
