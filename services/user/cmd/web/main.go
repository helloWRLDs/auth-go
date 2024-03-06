package main

import (
	"auth-go/pkg/drivers/mysql"
	"auth-go/services/user/configs"
	httpDelivery "auth-go/services/user/internal/delivery/http"
	"net/http"
	"os"
	"time"
)

func main() {
	cfg := configs.LoadAppConfig()
	db, err := mysql.GetMySqlConnection(cfg.Db)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()
	delivery := httpDelivery.NewRouter(db, cfg)
	delivery.InitRoutes()

	srv := &http.Server{
		Handler:      delivery.Router,
		Addr:         cfg.Addr,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
