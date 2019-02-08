package Data

import (
	"SimpleProject/Domin/Model"
	"database/sql"
)

func GetUserByUserName(UserName string) (*Model.User, error) {
	var user Model.User
	err := GetDB().QueryRowx(
		"app.GetUserByUserName",
		sql.Named("UserName", UserName),
	).StructScan(&user)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &user, nil

}
