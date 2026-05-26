package utils

import (
	"time"
	"todo-app/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

//generate token//

/*this is the signing secret*/

func GenerateJWT(userId int, cfg *config.Config) (string, error) {
	var jwtSecret = []byte(cfg.SECRETKEY)
	claims := jwt.MapClaims{
		"userID": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
