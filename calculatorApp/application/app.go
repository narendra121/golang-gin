package application

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartCalculator() {
	MapUrls()
	router.Run(":8080")
}
