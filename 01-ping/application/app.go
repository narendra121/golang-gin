package application

import (
	"github.com/gin-gonic/gin"
	"github.com/narendra121/golang-gin/01-ping/controller"
)

var (
	router = gin.Default()
)

func Startapplication() {
	router.GET("/hello", controller.Hello)
	router.Run(":8080")
}
