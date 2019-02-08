package Util

import (
	"SimpleProject/Domin/Model"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	validator "gopkg.in/go-playground/validator.v9"
)

var JwtKey = "6150645367566B5970337336763979244226452948404D6251655468576D5A71"

func HashAndSalt(password string) string {

	pwd := []byte(password)
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePasswords(hashedPwd string, password string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	plainPwd := []byte(password)

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}

func CreateTokenEndpoint(user *Model.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()
	claims := TokenClaims{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "bandzest-auth",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JwtKey))

	return tokenString, err
}

func ValidateToken(tokenString string) (bool, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})

	if err != nil {
		panic(err)
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}

func IsValid(model interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(model)
	return err
}

// func ValidateToken(tokenString string) (bool, error) {

// 	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected siging method")
// 		}
// 		return []byte(jwtKey), nil
// 	})

// 	log.Println(err)

// 	log.Println("Valid -- ", token.Valid)

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		_, err := userRepo.Find(claims["id"].(string))

// 		if err != nil {
// 			return false, err
// 		}

// 		return true, nil
// 	} else {
// 		return false, err
// 	}

// 	return false, nil
// }
