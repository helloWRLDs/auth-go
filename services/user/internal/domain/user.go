package domain

import (
	"regexp"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(email, password string) (*User, error) {
	newUser := &User{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	if err := newUser.IsValid(); err != nil {
		return nil, err
	}
	return newUser, nil
}

var (
	emailRegex    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	passwordRegex = `^[a-zA-Z\d]*[a-z][a-zA-Z\d]*[A-Z][a-zA-Z\d]*[a-zA-Z\d]*$`
)

func (u *User) IsValid() error {
	if len(u.Email) == 0 || len(u.Password) == 0 {
		return ErrEmptyField
	}
	if !regexp.MustCompile(emailRegex).MatchString(u.Email) {
		return ErrEmailValidation
	}
	if !regexp.MustCompile(passwordRegex).MatchString(u.Password) {
		return ErrPasswordValidation
	}
	return nil
}
