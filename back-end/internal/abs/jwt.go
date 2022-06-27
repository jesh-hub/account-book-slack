package abs

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func JwtDecode(credential string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(credential, nil)
	if token == nil {
		return nil, errors.New("cannot parse token")
	}
	//if !token.Valid {
	//	return nil, errors.New("token is unvalid")
	//}

	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}
