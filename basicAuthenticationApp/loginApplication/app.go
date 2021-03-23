package loginApplication

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartLogin() {
	MapUrls()
	router.Run(":8080")
}