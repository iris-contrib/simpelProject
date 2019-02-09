package Data

import (
	"SimpleProject/Domin/Dto"
	"SimpleProject/Domin/Model"
	"database/sql"
)

func CreateUser(user dto.User) (*int, error) {
	var userId int
	err := db.QueryRow("app.CreateUser",
		sql.Named("UserName", user.UserName),
		sql.Named("Password", user.Password),
		sql.Named("FirstName", user.FirstName),
		sql.Named("LastName", user.LastName),
		sql.Named("Email", user.Email),
	).Scan(&userId)

	if err != nil {
		return nil, err
	}
	return &userId, nil
}

func GetUser() (*[]Model.User, error) {
	var d []Model.User
	db.Select(&d, "app.GetUser")
	return &d, nil
}
