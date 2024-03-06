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
	d.Router.Group("/user")
	//	{
	//		d.Router.GET("/", d.UserUC.GetUsers)
	//		d.Router.GET("/:id", d.UserUC.GetUser)
	//		d.Router.PUT("/:id", d.UserUC.UpdateUser)
	//		d.Router.DELETE("/:id", d.UserUC.RemoveUser)
	//	}
	//}
}
