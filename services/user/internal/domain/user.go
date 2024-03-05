package domain

import (
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(email, password string) (*User, error) {
	newUser := &User{
		Email:     email,
		Password:  []byte(password),
		CreatedAt: time.Now(),
	}
	if err := newUser.IsValid(); err != nil {
		return nil, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(newUser.Password, 12)
	if err != nil {
		return nil, err
	}
	newUser.Password = hashedPassword
	return newUser, nil
}

var (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// passwordRegex = `^[a-zA-Z\d]*[a-z][a-zA-Z\d]*[A-Z][a-zA-Z\d]*[a-zA-Z\d]*$`
	passwordRegex = `[a-zA-Z0-9]{3,}`
)

func (u *User) IsValid() error {
	if len(u.Email) == 0 || len(u.Password) == 0 {
		return ErrEmptyField
	}
	if !regexp.MustCompile(emailRegex).MatchString(u.Email) {
		return ErrEmailValidation
	}
	if !regexp.MustCompile(passwordRegex).MatchString(string(u.Password)) {
		return ErrPasswordValidation
	}
	return nil
}
