package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var student Student

		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		context.Writer.Write([]byte("hello, " + student.Name))
	})

	engine.Run(":8090")
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}
