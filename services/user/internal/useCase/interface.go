package usecase

import "github.com/gin-gonic/gin"

type UserUseCase interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	RemoveUser(ctx *gin.Context)
}
