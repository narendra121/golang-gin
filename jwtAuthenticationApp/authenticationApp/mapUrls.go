package authenticationApp

import (
	"github.com/narendra121/golang-gin/jwtAuthenticationApp/controller/authController"
	"github.com/narendra121/golang-gin/jwtAuthenticationApp/controller/middleware"
)

func MapUrls() {
	router.POST("/login", authController.Login)
	router.GET("/welcome", middleware.AuthMiddleWare(),authController.Welcome)
}
