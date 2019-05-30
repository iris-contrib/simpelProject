package dto

import "github.com/dgrijalva/jwt-go"

// User dto ....
type User struct {
	FirstName string `json:"FirstName" validate:"required"`
	LastName  string `json:"LastName" validate:"required"`
	Email     string `json:"Email" validate:"required,email"`
	UserName  string `json:"UserName" validate:"required"`
	Password  string `json:"Password" validate:"required"`
}

// Login dto ....
type Login struct {
	UserName string `json:"UserName" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

//CreateUser dto ...
type CreateUser struct {
	UserID *int `json:"UserId"`
}

//NewCreateUser ...
func NewCreateUser(userID *int) *CreateUser {
	return &CreateUser{UserID: userID}
}

//TokenClaims ....
type TokenClaims struct {
	UserID    int32  `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	jwt.StandardClaims
}

//Token dto ...
type Token struct {
	AuthToken string `json:"Auth_Token"`
}

//NewToken ...
func NewToken(authToken string) *Token {
	return &Token{AuthToken: authToken}
}

//Response dto ....
type Response struct {
	Status       bool        `json:"Status"`
	Data         interface{} `json:"Data"`
	ErrorMessage string      `json:"ErrorMessage"`
}

//NewResponse ...
func NewResponse(status bool, data interface{}, errorMessage string) *Response {
	return &Response{Status: status, Data: data, ErrorMessage: errorMessage}
}
