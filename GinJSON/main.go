package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	engine := gin.Default()

	engine.GET("/hellojson", func(context *gin.Context) {

		fullPath := "请求路径: " + context.FullPath()
		fmt.Println(fullPath)

		resp := Response{Code: 1, Message: "Ok", Data: fullPath}
		/*
			context.JSON(200, map[string]interface{}{
				"code":    1,
				"message": "OK",
				"data":    fullPath,
			})
		*/
		context.JSON(200, &resp)
	})

	engine.Run(":8090")
}
