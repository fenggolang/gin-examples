package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义的html模板渲染
func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("./testdata/template/raw.tmpl")

	router.GET("/raw", func(context *gin.Context) {
		context.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2019, 12, 01, 0, 0, 0, 0, time.UTC),
		})
	})
	router.Run(":8080")
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	//return fmt.Sprintf("%d/%d/%d", year, month, day) // 2019/8/4
	return fmt.Sprintf("%d/%02d/%02d", year, month, day) // 2019/08/04 %02d表示变量占2位如果数字是0-9则，首位填充0。
}
