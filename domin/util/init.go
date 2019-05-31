package util

import "gopkg.in/go-playground/validator.v9"

var validate *validator.Validate

//InitValidator ...
func InitValidator() {
	validate = validator.New()
}
