package usecase

import (
	"auth-go/services/user/internal/domain"
	"auth-go/services/user/internal/repository"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserUseCaseImpl struct {
	repo repository.UserRepository
}

func NewUserUseCase(db *sql.DB) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		repo: repository.NewUserRepository(db),
	}
}

func (u *UserUseCaseImpl) RegisterUser(ctx *gin.Context) {
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.JSON(422, err.Error())
		return
	}
	email := ctx.Request.PostFormValue("email")
	password := ctx.PostForm("password")
	fmt.Println(email, password)

	user, err := domain.NewUser(email, password)
	if err != nil {
		ctx.JSON(422, err.Error())
		return
	}

	ctx.JSON(200, user)
}
