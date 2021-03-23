package application

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartAuthApp() {
	MapUrls()
	router.Run(":8080")
}