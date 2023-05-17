package server

import (
	"github.com/rs/cors"
	"net/http"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"post"`
	Salt string `json:"salt"`
}

func InitServer(cfg Config) error {
	r := initRoute()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	return http.ListenAndServe(
		cfg.Host+":"+cfg.Port,
		c.Handler(r),
	)
}
