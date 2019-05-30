package util

import "github.com/dgrijalva/jwt-go"

type tokenClaims struct {
	UserID    int32  `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	jwt.StandardClaims
}
