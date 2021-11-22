package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.DELETE("/user/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		userID := context.Param("id")

		fmt.Println(userID)

		context.Writer.Write([]byte("Delete user's id: " + userID))
	})

	engine.Run(":8090")
}
