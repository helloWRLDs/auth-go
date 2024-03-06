package httpDelivery

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JsonContentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}

func SecureHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.Writer.Header().Set("X-Frame-Options", "deny")
		ctx.Writer.Header().Set("X-XSS-Protection", "0")
		ctx.Next()
	}
}

func EncodedContentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		ctx.Next()
	}
}

var secretKey = "your_secret_key_here"

type UserClaims struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  []byte    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	jwt.StandardClaims
}

func getSecretKey() string {
	return secretKey
}



func AuthMiddleware(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed (no token)"})
		ctx.Abort()
		return
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(getSecretKey()), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed (invalid token)"})
		ctx.Abort()
		return
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		ctx.Set("user", claims)
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed (invalid token)"})
		ctx.Abort()
		return
	}
}
