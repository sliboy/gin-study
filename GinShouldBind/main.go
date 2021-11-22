package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Register struct {
	UserName string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"pwd"`
}

func main() {
	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var _register Register

		err := context.ShouldBind(&_register)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(_register.UserName)
		fmt.Println(_register.Phone)

		context.Writer.Write([]byte(_register.UserName + " Register"))
	})

	engine.Run(":8090")
}
