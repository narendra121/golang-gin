package application

import "github.com/narendra121/golang-gin/hello/controller"

func MapUrls() {
	router.GET("/ping", controller.Hello)
}