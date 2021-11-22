package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		context.Writer.WriteString(fullPath)
	})

	engine.Run(":8090")
}
