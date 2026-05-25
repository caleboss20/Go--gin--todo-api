package services

import (
	"errors"
	"todo-app/internal/models"
	"todo-app/internal/repository"
	"todo-app/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(loginInput models.LoginInput) (string, error) {

	//find user by email//
	user, err := repository.GetUserByEmail(loginInput.Email)
	if err != nil {
		return "", errors.New("invaid email or password")
	}
	//check if password is same as hashed password in db//
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(loginInput.Password))

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	//create a token//
	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	//send token back to user//
	return token, err
}
