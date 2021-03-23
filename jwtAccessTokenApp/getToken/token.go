package getToken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/narendra121/golang-gin/jwtAccessTokenApp/domain/user"
)

var(
	tokenSecret=[]byte("hellosecret")
)
func Token(user user.User) (string, error) {
		claims:=jwt.MapClaims{}
		claims["user_email"]=user.Email
		claims["exp"]=time.Now().Add(time.Minute*5).Unix()
		token:=jwt.NewWithClaims(jwt.SigningMethodHS256,&claims)
		tokenSecret,err:=token.SignedString(tokenSecret)
		return tokenSecret,err
}