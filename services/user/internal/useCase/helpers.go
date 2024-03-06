package usecase

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(userId int, userEmail string) (string, error) {
	claim := &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]string{
			"id":    string(userId),
			"email": userEmail,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(os.Getenv("secret-key")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
