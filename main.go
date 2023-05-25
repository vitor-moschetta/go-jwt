package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func main() {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["audioId"] = "60748899"
	claims["ldap"] = "vmoschetta"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	secret := ""
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	println(tokenString)
}
