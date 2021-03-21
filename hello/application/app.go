package application

import (
	"github.com/gin-gonic/gin"
	"github.com/narendra121/golang-gin/hello/controller"
)

var (
	//Returns engine
	router = gin.Default()
)

func Startapplication() {
	//Url: localhost:8080/hello
	router.GET("/hello", controller.Hello)
	router.Run(":8080")
}
