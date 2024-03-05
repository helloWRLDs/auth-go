package httpDelivery

import (
	usecase "auth-go/services/user/internal/useCase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type HttpDelivery struct {
	Router *gin.Engine
	UserUC usecase.UserUseCase
}

func NewRouter(db *sql.DB) *HttpDelivery {
	return &HttpDelivery{
		Router: gin.Default(),
		UserUC: usecase.NewUserUseCase(db),
	}
}

func (d *HttpDelivery) InitRoutes() {
	d.Router.Use(JsonContentMiddleware(), SecureHeaders())
	d.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})

	d.Router.POST("/signup", d.UserUC.RegisterUser)
}
