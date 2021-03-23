package authController

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/narendra121/golang-gin/jwtAuthenticationApp/authentication/tokenGen"
	"github.com/narendra121/golang-gin/jwtAuthenticationApp/domain/user"
)

func Login(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json body"})
		return
	}
	token, tErr := tokenGen.Token(user)
	if tErr != nil {
		c.JSON(http.StatusBadRequest, tErr)
		return
	}
	tokenInfo := map[string]string{
		"name":  user.Name,
		"token": token,
	}
	c.JSON(http.StatusOK, gin.H{"LoginData": tokenInfo})

}

func Welcome(c *gin.Context) {

	c.JSON(http.StatusOK, "Login Successful")
}
