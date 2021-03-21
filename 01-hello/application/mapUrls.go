package application

import "github.com/narendra121/golang-gin/01-hello/controller"

func MapUrls() {
	router.GET("/ping", controller.Hello)
}