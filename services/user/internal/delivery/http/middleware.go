package httpDelivery

import "github.com/gin-gonic/gin"

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
