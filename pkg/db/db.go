package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db"`
}

func InitDB(cfg Config) (*sqlx.DB, error) {
	return sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=disable",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.DBName,
		),
	)
}
