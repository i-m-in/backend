package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/melnk300/CalendarBack/internal/models"
	"github.com/melnk300/CalendarBack/internal/server"
	"github.com/melnk300/CalendarBack/pkg/db"
	"github.com/melnk300/CalendarBack/utils"
	"log"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// INFO: init Config
	database, _ := os.LookupEnv("POSTGRES_DATABASE")
	user, _ := os.LookupEnv("POSTGRES_USER")
	password, _ := os.LookupEnv("POSTGRES_PASSWORD")
	stage, _ := os.LookupEnv("STAGE")
	salt, _ := os.LookupEnv("SALT")

	dbCfg := db.Config{
		Host:     "127.0.0.1",
		User:     user,
		Password: password,
		DBName:   database,
	}
	srvCfg := server.Config{
		Host: "",
		Port: "8080",
	}

	_ = utils.Config{
		Salt: salt,
	}

	// INFO: init DB
	db, err := db.InitDB(dbCfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// INFO: init models (Domain model)
	models.SetDB(db)

	// INFO: init server
	if err := server.InitServer(srvCfg); err != nil {
		log.Panic(err)
	}

	// make migrations from sql files
	if stage == "dev" {
		fmt.Println("Make stage=PROD if you want to use production mode")
		models.MakeMigrations()
		fmt.Println("Migrations done")
	}
}
