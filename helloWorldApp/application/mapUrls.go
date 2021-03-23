package application

import "github.com/narendra121/golang-gin/helloWorldApp/controller"

func MapUrls() {
	router.GET("/ping", controller.Hello)
}