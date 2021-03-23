package tokenGen

import (
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/narendra121/golang-gin/jwtAuthenticationApp/domain/user"
)

var(
	TokenSecret=[]byte("hellosecret")
)
func Token(user user.User) (string, error) {
		claims:=jwt.MapClaims{}
		claims["user_email"]=user.Email
		claims["exp"]=time.Now().Add(time.Minute*5).Unix()
		token:=jwt.NewWithClaims(jwt.SigningMethodHS256,&claims)
		authToken,err:=token.SignedString(TokenSecret)
		return authToken,err
}


func TokenValid(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("token sign method error %v", token.Header["alg"])
		}
		return TokenSecret, nil
	})
	if err != nil {
		fmt.Printf("Err  --- %v \n", err)
		return false, ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userEmail := claims["user_email"]
		return true, userEmail.(string)
	} else {
		fmt.Printf("The alg header %v \n", claims["alg"])
		return false, "user_email"
	}
}
