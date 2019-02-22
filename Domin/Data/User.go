package data

import (
	"SimpleProject/Domin/Dto"
	"SimpleProject/Domin/model"
	"database/sql"
)

//CreateUser ...
func CreateUser(user dto.User) (*int, error) {
	var UserID int
	err := db.QueryRow("app.CreateUser",
		sql.Named("UserName", user.UserName),
		sql.Named("Password", user.Password),
		sql.Named("FirstName", user.FirstName),
		sql.Named("LastName", user.LastName),
		sql.Named("Email", user.Email),
	).Scan(&UserID)

	if err != nil {
		return nil, err
	}
	return &UserID, nil
}

// Get All User ...
func Get() (*[]model.User, error) {
	var d []model.User
	err := db.Select(&d, "app.GetUser")
	return &d, err
}

//GetUser : Get User By UserId  ...
func GetUser(userID float64) (*model.User, error) {

	var user model.User
	err := db.QueryRowx(
		"app.GetUserByUserId",
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
		"app.GetUserByUserName",
		sql.Named("UserName", UserName),
	).StructScan(&user)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &user, nil

}
