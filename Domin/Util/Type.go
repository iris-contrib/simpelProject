package Util

import jwt "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserId    uint64 `json:"UserId"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	jwt.StandardClaims
}
