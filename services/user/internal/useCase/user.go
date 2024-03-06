package usecase

import (
	"auth-go/services/user/internal/domain"
	"auth-go/services/user/internal/repository"
	"database/sql"
	"fmt"

	"strconv"

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

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	user, err := domain.NewUser(email, password)
	if err != nil {
		ctx.JSON(422, err.Error())
		return
	}

	if u.repo.ExistsByEmail(user.Email) {
		ctx.JSON(409, "user with such email already exists")
		return
	}

	id, err := u.repo.Insert(*user)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, fmt.Sprintf("user registered with id=%d", id))
}

func (u *UserUseCaseImpl) LoginUser(ctx *gin.Context) {

	if err := ctx.Request.ParseForm(); err != nil {
		ctx.JSON(422, err.Error())
		return
	}
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if !u.repo.ExistsByEmail(email) {
		ctx.JSON(401, fmt.Sprintf("Not authorized: email=%s not found", email))
		return
	}

	userID, err := u.repo.Authenticate(email, []byte(password))
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Not authorized: %s", err.Error()))
		return
	}

	token, err := generateToken(userID, email)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Not authorized: %s", err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}

func (u *UserUseCaseImpl) GetUsers(ctx *gin.Context) {
	users, err := u.repo.GetAll()
	if err != nil {
		ctx.JSON(500, "Internal Server Error")
		return
	}
	ctx.JSON(200, users)
}

func (u *UserUseCaseImpl) GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := u.repo.Get(id)
	if err != nil {

		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, user)
}

func (u *UserUseCaseImpl) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := ctx.Request.ParseForm(); err != nil {
		ctx.JSON(422, err.Error())
		return
	}
	var userToUpdate domain.User

	userToUpdate.Email = ctx.PostForm("email")
	userToUpdate.Password = []byte(ctx.PostForm("password"))

	existingUser, err := u.repo.Get(id)
	if err != nil {
		ctx.JSON(404, "user doesn't exist")
		return
	}
	if err := userToUpdate.ValidateEmail(); err == nil {
		existingUser.Email = userToUpdate.Email
	}

	if err := userToUpdate.ValidatePassword(); err == nil {
		existingUser.SetPassword(userToUpdate.Password)
	}

	if err := u.repo.Update(id, existingUser); err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User updated successfully"})
}

func (u *UserUseCaseImpl) RemoveUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := u.repo.Delete(id); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
