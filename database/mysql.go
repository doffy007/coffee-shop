package database

import (
	"fmt"
	"log"
	"time"

	"github.com/doffy007/coffee-shop/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Mysql() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load file .env")
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		config.ConfigureApp.DBUsername,
		config.ConfigureApp.DBPassword,
		config.ConfigureApp.DBHost,
		config.ConfigureApp.DBName,
	)

	db, err := sqlx.Open("mysql", conn)
	if err != nil {
		log.Fatal("Error connection")
	}

	db.SetConnMaxIdleTime(time.Duration(config.ConfigureApp.DBConnectionIdle) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(config.ConfigureApp.DBConnectionLifetime) * time.Minute)
	db.SetMaxIdleConns(config.ConfigureApp.DBMaxIdle)
	db.SetMaxOpenConns(config.ConfigureApp.DBMaxOpen)

	return db
}
