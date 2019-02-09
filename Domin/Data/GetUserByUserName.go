package data

import (
	"SimpleProject/Domin/model"
	"database/sql"
)

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
