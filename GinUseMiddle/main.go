package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求Path: ", path)
		fmt.Println("请求Method: ", method)
		context.Next()
		fmt.Println(context.Writer.Status())
	}
}

func main() {
	engine := gin.Default()
	engine.Use(RequestInfos())

	engine.GET("/query", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})
	userGroup := engine.Group("/user")
	userGroup.GET("/register", registerHandle)
	userGroup.GET("/login", loginHandle)
	userGroup.GET("/info", infoHandle)

	engine.Run(":9000")
}

func registerHandle(context *gin.Context) {
	fullPath := " 用户注册功能 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func loginHandle(context *gin.Context) {
	fullPath := " 用户登录功能 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func infoHandle(context *gin.Context) {
	fullPath := " 信息查看功能 " + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}
