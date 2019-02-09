package model

//User model ....
type User struct {
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`
	Email     string `db:"Email"`
	UserName  string `db:"UserName"`
	Password  string `db:"Password"`
	RoleID    uint8  `db:"RoleId"`
	UserID    uint64 `db:"UserId"`
}
