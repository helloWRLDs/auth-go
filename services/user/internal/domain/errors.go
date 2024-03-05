package domain

import "errors"

var (
	ErrEmptyField         = errors.New("fields cannot be empty(password, email)")
	ErrEmailValidation    = errors.New("wrong email format")
	ErrPasswordValidation = errors.New("password must have atleast 1 number, 1 uppercase and lowercase letters")
)
