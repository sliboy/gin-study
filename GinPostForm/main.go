package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println(username)
		}

		password, exists := context.GetPostForm("pwd")
		if exists {
			fmt.Println(password)
		}

		context.Writer.Write([]byte("Hello , " + username))
	})

	engine.Run(":8090")
}
