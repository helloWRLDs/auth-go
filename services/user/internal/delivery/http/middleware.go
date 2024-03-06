package httpDelivery

import (
	"fmt"
	"net/http"
	"os"

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

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqToken := ctx.Request.Header.Get("Authorization")
		if reqToken == "" {
			ctx.JSON(http.StatusUnauthorized, "Authorization Failed (no token)")
			ctx.Abort()
			return
		}
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("secret_key")), nil
		})
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusUnauthorized, "Authorization Failed (invalid token)")
			ctx.Abort()
			return
		}
		if err := claims.Valid(); err != nil {
			ctx.JSON(http.StatusUnauthorized, "Authorization Failed (expired)")
			ctx.Abort()
			return
		}
		data, ok := claims["data"].(map[string]interface{})
		if !ok {
			ctx.JSON(http.StatusUnauthorized, "Authorization Failed (invalid data)")
			ctx.Abort()
			return
		}
		ctx.Set("email", data["email"])
		ctx.Set("id", data["id"])
		ctx.Next()
	}
}

// func AuthMiddleware(ctx *gin.Context) {
// 	tokenString := ctx.Request.Header.Get("Authorization")
// 	if tokenString == "" {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed (no token)"})
// 		ctx.Abort()
// 		return
// 	}
// 	token, err := jwt.ParseWithClaims(tokenStriserClaing, &Ums{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("secret_key")), nil
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed (invalid token)"})
// 		ctx.Abort()
// 		return
// 	}
// 	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
// 		ctx.Set("user", claims)
// 		ctx.Next()
// 	} else {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization Failed (invalid token)"})
// 		ctx.Abort()
// 		return
// 	}
// }
