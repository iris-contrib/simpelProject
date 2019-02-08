package dto

type User struct {
	FirstName string `json:"FirstName" validate:"required"`
	LastName  string `json:"LastName" validate:"required"`
	Email     string `json:"Email" validate:"required,email"`
	UserName  string `json:"UserName" validate:"required"`
	Password  string `json:"Password" validate:"required"`
}

type Login struct {
	UserName string `json:"UserName" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

type CreateUser struct {
	UserId *int `json:"UserId"`
}
