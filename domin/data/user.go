package data

import (
	"database/sql"
	"fmt"

	"github.com/majidbigdeli/simpelProject/domin/dto"
	"github.com/majidbigdeli/simpelProject/domin/model"
)

//CreateUser ...
func CreateUser(user dto.User) (*int, error) {
	var userID int
	err := db.QueryRow("dbo.CreateUser",
		sql.Named("FirstName", user.FirstName),
		sql.Named("LastName", user.LastName),
		sql.Named("Email", user.Email),
		sql.Named("UserName", user.UserName),
		sql.Named("Password", user.Password),
		sql.Named("Mobile", user.Mobile),
		sql.Named("NationalCode", user.NationalCode),
	).Scan(&userID)

	if err != nil {
		return nil, err
	}

	if userID == 0 {
		return nil, fmt.Errorf("username is exist")
	}

	return &userID, nil
}

// Get All User ...
func Get() (*[]model.User, error) {
	var d []model.User
	err := db.Select(&d, "dbo.GetUser")
	return &d, err
}

//GetUser : Get User By UserId  ...
func GetUser(userID float64) (*model.User, error) {

	var user model.User
	err := db.QueryRowx(
		"dbo.GetUserByUserId",
		sql.Named("UserId", userID),
	).StructScan(&user)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &user, nil

}

//GetUserByUserName ..
func GetUserByUserName(UserName string) (*model.User, error) {
	var user model.User
	err := db.QueryRowx(
		"dbo.GetUserByUserName",
		sql.Named("UserName", UserName),
	).StructScan(&user)

	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("username not exist")
	}

	return &user, nil

}
