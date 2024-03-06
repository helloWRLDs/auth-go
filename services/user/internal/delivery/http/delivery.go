package httpDelivery

import (
	"auth-go/services/user/configs"
	usecase "auth-go/services/user/internal/useCase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type HttpDelivery struct {
	Router    *gin.Engine
	UserUC    usecase.UserUseCase
	AppConfig *configs.AppConfig
}

func NewRouter(db *sql.DB, cfg *configs.AppConfig) *HttpDelivery {
	return &HttpDelivery{
		Router:    gin.Default(),
		UserUC:    usecase.NewUserUseCase(db),
		AppConfig: cfg,
	}
}

func (d *HttpDelivery) InitRoutes() {
	d.Router.Use(JsonContentMiddleware(), SecureHeaders())
	d.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})
	d.Router.POST("/signup", d.UserUC.RegisterUser)
	d.Router.POST("/login", d.UserUC.LoginUser)
	userGroup := d.Router.Group("/users", Authenticate())
	{
		userGroup.GET("/", d.UserUC.GetUsers)
		userGroup.GET("/:id", d.UserUC.GetUser)
		userGroup.PUT("/:id", d.UserUC.UpdateUser)
		userGroup.DELETE("/:id", d.UserUC.RemoveUser)
	}
}
