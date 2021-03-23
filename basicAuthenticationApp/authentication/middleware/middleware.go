package middleware

import "github.com/gin-gonic/gin"

func BasicAuthentication() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{"narendra": "123456"})
}
