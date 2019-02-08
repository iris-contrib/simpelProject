package Model

type User struct {
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`
	Email     string `db:"Email"`
	UserName  string `db:"UserName"`
	Password  string `db:"Password"`
	RoleId    uint8  `db:"RoleId"`
	UserId    uint64 `db:"UserId"`
}
