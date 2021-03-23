package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/narendra121/golang-gin/jwtAuthenticationApp/authentication/tokenGen"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")

		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authenti"})
			c.Abort()
			return
		}
		token := split[1]
		valid, userEmail := tokenGen.TokenValid(token)

		if valid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authentic"})
			c.Abort()
		} else {
			c.Set("user_email", userEmail)
			c.Next()
		}

	}

}
