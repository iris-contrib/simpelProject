package dto

// User dto ....
type User struct {
	FirstName string `json:"FirstName" validate:"required"`
	LastName  string `json:"LastName" validate:"required"`
	Email     string `json:"Email" validate:"required,email"`
	UserName  string `json:"UserName" validate:"required"`
	Password  string `json:"Password" validate:"required"`
}

// Login dto ....
type Login struct {
	UserName string `json:"UserName" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

//CreateUser dto ...
type CreateUser struct {
	UserID *int `json:"UserId"`
}
