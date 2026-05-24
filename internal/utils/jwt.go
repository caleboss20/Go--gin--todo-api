package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//generate token//

/*this is the signing secret*/
var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(userId int) (string, error) {
	claims := jwt.MapClaims{
		"userID": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
