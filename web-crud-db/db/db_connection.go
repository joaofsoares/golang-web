package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/Valgard/godotenv"
	"github.com/go-sql-driver/mysql"
)

func createConfig() mysql.Config {

	godotenv.Load(".env")

	return mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    os.Getenv("DB_NET"),
		Addr:   os.Getenv("DB_ADDRESS"),
		DBName: os.Getenv("DB_SCHEMA"),
	}
}

func createConnection() *sql.DB {
	cfg := createConfig()

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal("DB Connection Failed")
	}

	return db
}
