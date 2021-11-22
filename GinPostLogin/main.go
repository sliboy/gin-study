package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		username := context.PostForm("username")
		fmt.Println(username)

		password := context.PostForm("pwd")
		fmt.Println(password)

		context.Writer.Write([]byte("User login"))
	})
	engine.Run(":8090")
}
