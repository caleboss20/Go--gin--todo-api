package utils

import (
	"errors"
	"todo-app/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenstring string, cfg *config.Config) (int, error) {
	//jwtsecret being loaded here from config
	jwtSecret := []byte(cfg.SECRETKEY)

	//tokenstring ="ehf2gu2f433.... sent by client in header"//
	//middleware will call this function and pass the token here to verify//
	//jwt.parse opens the token string and breaks it into header,payload,signature//

	//parse the token string and verify it using our secret key//
	//during parsing jwt.parse calls this callback func asking "what key should i use? "//
	token, err := jwt.Parse(tokenstring,
		func(token *jwt.Token) (interface{}, error) {

			//we check the algorithn is HS256 (HMAC family )//
			//prevents:None Algorithm Attack + Algorithm Confusion Attack//
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected signing method")
			}

			//return our secret key so jwt.parse can verify the signature//
			//prevents:Token Forgery Attack + Payload Tampering Attack
			return jwtSecret, nil //this is the key it tell it to use//

		})
	//jwt.Parse failed meaning token is invalid ,tampered,or expired
	//Prevents:Expired Token Reuse Attack+Signature Stripping attack//

	if err != nil {
		return 0, errors.New("invalid token")
	}
	// cast the claims inside the token to MapClaims (key value pairs)
	// MapClaims is basically map[string]interface{}
	// token.Valid checks token is not expired or tampered
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}
	// read the userID we stored inside the token in GenerateJWT
	// userID comes back as float64 because JWT stores all numbers as JSON numbers internally
	// PREVENTS: Claims Injection Attack (we only read what we stored)
	userId, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// convert float64 back to int and return userId to middleware
	// middleware will use this userId to identify who made the request

	return int(userId), nil
	//even though userid is converted to int in go interface{} is still has no type

}
