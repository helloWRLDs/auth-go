package main

import (
	"auth-go/services/user/configs"
	"fmt"
	"net/http"
	"time"
)

func main() {
	cfg := configs.LoadAppConfig()
	srv := &http.Server{
		Addr:         cfg.Addr,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	fmt.Println(srv)
}
