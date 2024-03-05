package usecase

import "github.com/gin-gonic/gin"

type UserUseCase interface {
	RegisterUser(ctx *gin.Context)
}
