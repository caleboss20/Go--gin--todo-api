package models

//for the user register struct //
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//for the user login struct //

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
