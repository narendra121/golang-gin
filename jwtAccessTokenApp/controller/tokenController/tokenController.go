package tokenController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narendra121/golang-gin/jwtAccessTokenApp/domain/user"
	"github.com/narendra121/golang-gin/jwtAccessTokenApp/getToken"
)

func Token(c *gin.Context) {
	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Json Body ...")
		return
	}
	token, terr := getToken.Token(user)
	if terr != nil {
		c.JSON(http.StatusInternalServerError, terr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Access Token": token})
}
