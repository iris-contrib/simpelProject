package data

import (
	"SimpleProject/Domin/Dto"
	"SimpleProject/Domin/model"
	"database/sql"
)

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

func GetUser() (*[]model.User, error) {
	var d []model.User
	db.Select(&d, "app.GetUser")
	return &d, nil
}
