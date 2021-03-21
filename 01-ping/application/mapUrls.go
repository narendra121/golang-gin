package application

import "github.com/narendra121/golang-gin/01-ping/controller"

func MapUrls() {
	router.GET("/ping", controller.Hello)
}