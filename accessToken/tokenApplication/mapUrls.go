package tokenApplication

import "github.com/narendra121/golang-gin/accessToken/controller/tokenController"

func MapUrls() {
	router.POST("/token", tokenController.Token)
}