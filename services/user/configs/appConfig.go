package configs

import (
	"auth-go/pkg/drivers/mysql"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Addr      string
	SecretKey string
	Db        *mysql.MySqlConfig
}

func LoadAppConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		return nil
	}
	return &AppConfig{
		Addr:      os.Getenv("addr"),
		SecretKey: os.Getenv("secret_key"),
		Db: &mysql.MySqlConfig{
			User:     os.Getenv("db_user"),
			Password: os.Getenv("db_password"),
			Name:     os.Getenv("db_name"),
		},
	}
}
