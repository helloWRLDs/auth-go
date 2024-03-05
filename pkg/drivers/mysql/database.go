package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlConfig struct {
	User     string
	Password string
	Name     string
}

func GetMySqlConnection(cfg *MySqlConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@/%s?parseTime=true", cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
