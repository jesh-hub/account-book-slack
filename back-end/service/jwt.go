package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func JwtDecode(credential string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(credential, nil)
	if token == nil {
		return nil, errors.New("cannot parse token")
	}
	//if !token.Valid {
	//	return nil, errors.New("token is invalid")
	//}

	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}

func CreateJwtToken(userId string, email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = userId
	atClaims["email"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
