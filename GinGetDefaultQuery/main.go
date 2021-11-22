package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		//		name := context.DefaultQuery("name", "default")
		name := context.Query("name")
		fmt.Println(name)

		context.Writer.Write([]byte("Hello , " + name))
	})

	err := engine.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}
}
