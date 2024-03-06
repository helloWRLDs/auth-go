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
	// receiving data from post request
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	// validating and creating user's instance
	user, err := domain.NewUser(email, password)
	if err != nil {
		ctx.JSON(422, err.Error())
		return
	}
	// check for uniqueness of email
	if u.repo.ExistsByEmail(user.Email) {
		ctx.JSON(409, "user with such email already exists")
		return
	}
	// inserting new User to database
	id, err := u.repo.Insert(*user)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, fmt.Sprintf("user registered with id=%d", id))
}
func (u *UserUseCaseImpl) LoginUser(ctx *gin.Context) {
	var userRepository repository.UserRepositoryImpl
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.JSON(422, err.Error())

	}

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if u.repo.ExistsByEmail(email) {
		userID, err := userRepository.Authenticate(email, password)
		if err != nil {
			ctx.JSON(401, fmt.Sprintf("Not authorized"))
			return
		}
		ctx.JSON(200, gin.H{"message": "Login successful", "userID": userID})
	} else {
		ctx.JSON(401, fmt.Sprintf("email=%s not found", email))
	}

}
