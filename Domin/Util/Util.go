package util

import (
	"SimpleProject/Domin/model"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/context"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

//JwtKey ...
var jwtKey = "6150645367566B5970337336763979244226452948404D6251655468576D5A71"

//MyJwtMiddleware ...
var MyJwtMiddleware = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

//HashAndSalt ...
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

//ComparePasswords ...
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

//CreateTokenEndpoint ...
func CreateTokenEndpoint(user *model.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()
	claims := tokenClaims{
		UserID:    user.UserID,
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

	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}

//ValidateToken ...
//func ValidateToken(tokenString string) (bool, error) {
//
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte(jwtKey), nil
//	})
//
//	if err != nil {
//		panic(err)
//	}
//
//	if !token.Valid {
//		return false, nil
//	}
//
//	return true, nil
//}

//IsValid ...
func IsValid(model interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(model)
	return err
}

//GetUserIDFromToken ....
func GetUserIDFromToken(ctx context.Context) float64 {
	userToken := MyJwtMiddleware.Get(ctx)

	var userID float64

	if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
		userID = claims["UserID"].(float64)
	}

	return userID
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
