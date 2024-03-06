package repository

import "errors"

var (
	ErrIncorrectPassword = errors.New("password is incorrect")
	ErrEmailNotFound     = errors.New("email not found")
)
