package util

import jwt "github.com/dgrijalva/jwt-go"

type tokenClaims struct {
	UserID    uint64 `json:"UserID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	jwt.StandardClaims
}
