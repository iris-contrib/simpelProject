package dto

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// UserDto dto ....
type UserDto struct {
	FirstName    string `json:"FirstName" validate:"required"`
	LastName     string `json:"LastName" validate:"required"`
	Email        string `json:"Email" validate:"required,email"`
	UserName     string `json:"UserName" validate:"required"`
	Password     string `json:"Password" validate:"required"`
	Mobile       string `json:"Mobile" validate:"required"`
	NationalCode string `json:"NationalCode" validate:"required"`
}

// LoginDto dto ....
type LoginDto struct {
	UserName string `json:"UserName" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

//ReturnID dto ...
type ReturnID struct {
	UserID *int `json:"UserId"`
}

//TokenClaims ....
type TokenClaims struct {
	UserID    int32  `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	jwt.StandardClaims
}

//TokenDto dto ...
type TokenDto struct {
	AuthToken string `json:"Auth_Token"`
}

type customError struct {
	Message string
}

//ResponseDto dto ....
type ResponseDto struct {
	Status       bool         `json:"Status"`
	Data         interface{}  `json:"Data"`
	ErrorMessage *customError `json:"ErrorMessage"`
}

//NewResponse ...
func NewResponse(status bool, data interface{}, errorMessage error) *ResponseDto {

	if errorMessage != nil {
		return &ResponseDto{Status: status, Data: data, ErrorMessage: &customError{Message: errorMessage.Error()}}
	}

	return &ResponseDto{Status: status, Data: data, ErrorMessage: nil}
}

//NewReturnID ...
func NewReturnID(userID *int) *ReturnID {
	return &ReturnID{UserID: userID}
}

//NewToken ...
func NewToken(authToken string) *TokenDto {
	return &TokenDto{AuthToken: authToken}
}

func (e *customError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
