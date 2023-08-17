package db

import (
	"database/sql"
	"github.com/Valgard/godotenv"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"web-crud-db/model"
)

func loadProperties() model.DbConfig {

	godotenv.Load(".env")

	return model.DbConfig{
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASS"),
		Net:     os.Getenv("DB_NET"),
		Address: os.Getenv("DB_ADDRESS"),
		DbName:  os.Getenv("DB_SCHEMA"),
	}
}

func createConfig() mysql.Config {

	dbConfig := loadProperties()

	return mysql.Config{
		User:   dbConfig.User,
		Passwd: dbConfig.Pass,
		Net:    dbConfig.Net,
		Addr:   dbConfig.Address,
		DBName: dbConfig.DbName,
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
