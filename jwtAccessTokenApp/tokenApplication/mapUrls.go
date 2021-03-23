package tokenApplication

import "github.com/narendra121/golang-gin/jwtAccessTokenApp/controller/tokenController"

func MapUrls() {
	router.POST("/token", tokenController.Token)
}