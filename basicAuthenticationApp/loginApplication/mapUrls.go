package loginApplication

import (
	"github.com/narendra121/golang-gin/basicAuthenticationApp/authentication/middleware"
	"github.com/narendra121/golang-gin/basicAuthenticationApp/controller/loginController"
)


func MapUrls() {

	router.GET("/login",middleware.BasicAuthentication(),loginController.Login)
}

