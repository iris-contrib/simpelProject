package dto

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// User dto ....
type User struct {
	FirstName    string `json:"FirstName" validate:"required"`
	LastName     string `json:"LastName" validate:"required"`
	Email        string `json:"Email" validate:"required,email"`
	UserName     string `json:"UserName" validate:"required"`
	Password     string `json:"Password" validate:"required"`
	Mobile       string `json:"Mobile" validate:"required"`
	NationalCode string `json:"NationalCode" validate:"required"`
}

// Login dto ....
type Login struct {
	UserName string `json:"UserName" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

//ReturnID dto ...
type ReturnID struct {
	UserID *int `json:"UserId"`
}

//NewReturnID ...
func NewReturnID(userID *int) *ReturnID {
	return &ReturnID{UserID: userID}
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

type argError struct {
	Message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

//Response dto ....
type Response struct {
	Status       bool        `json:"Status"`
	Data         interface{} `json:"Data"`
	ErrorMessage *argError   `json:"ErrorMessage"`
}

//NewResponse ...
func NewResponse(status bool, data interface{}, errorMessage error) *Response {

	if errorMessage != nil {
		return &Response{Status: status, Data: data, ErrorMessage: &argError{Message: errorMessage.Error()}}
	}
	return &Response{Status: status, Data: data, ErrorMessage: nil}
}
