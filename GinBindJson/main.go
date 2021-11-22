package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  string `form:"age"`
}

func main() {
	engine := gin.Default()

	engine.POST("/addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var person Person

		err := context.BindJSON(&person)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(person.Name)
		fmt.Println(person.Age)

		context.Writer.Write([]byte("添加记录：" + person.Name))
	})

	engine.Run(":8090")
}
