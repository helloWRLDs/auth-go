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
	// receiving data from post request
	if err := ctx.Request.ParseForm(); err != nil {
		ctx.JSON(422, err.Error())
		return
	}
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	// authenticating email
	if !u.repo.ExistsByEmail(email) {
		ctx.JSON(401, fmt.Sprintf("Not authorized: email=%s not found", email))
		return
	}
	// authenticating password
	userID, err := u.repo.Authenticate(email, []byte(password))
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Not authorized: %s", err.Error()))
		return
	}
	// signing token
	token, err := generateToken(userID, email)
	if err != nil {
		ctx.JSON(401, fmt.Sprintf("Not authorized: %s", err.Error()))
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}

func (u *UserUseCaseImpl) GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	fmt.Println(userID)
	id, err := strconv.Atoi(userID)
	if err != nil {

		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve user from the repository
	user, err := u.repo.Get(id)
	if err != nil {
		// Handle error
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, user)
}

func (u *UserUseCaseImpl) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		// Handle error if the user ID is not a valid integer
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve user data from the request body
	var userData domain.User
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// Update user
	if err := u.repo.Update(id, userData); err != nil {
		// Handle error (e.g., user not found)
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

	// Delete user from the repository
	if err := u.repo.Delete(id); err != nil {
		// Handle error
		ctx.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
