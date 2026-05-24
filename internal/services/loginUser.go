package services

import (
	"errors"
	"todo-app/internal/models"
	"todo-app/internal/repository"
	"todo-app/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(loginInput models.LoginInput) (string, error) {

	//look up requested user//
	user, err := repository.GetUserByEmail(loginInput.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	//compare sent in password with saved user password hash//
	//bring the hash password from db first,order matters//
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	//generate jwt token//
	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return "", err
	}

	//==========send back to user===to login handler=====//

	return token, nil

}
