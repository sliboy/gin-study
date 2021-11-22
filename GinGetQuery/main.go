package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		username := context.Query("name")
		fmt.Println(username)

		context.Writer.Write([]byte("Hello " + username))
	})

	engine.Run(":8090")
}
