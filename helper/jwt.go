package helper

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/kristain09/API4/config"
)

func GenerateJWT(id string) string {
	var informasi = jwt.MapClaims{}
	informasi["id"] = id

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, informasi)

	resultToken, err := rawToken.SignedString([]byte(config.Secret))
	if err != nil {
		log.Println("generate jwt error ", err.Error())
		return ""
	}

	return resultToken
}

func DecodeJWT(token *jwt.Token) string {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		user_id := data["id"].(string)

		return user_id
	}

	return ""
}
