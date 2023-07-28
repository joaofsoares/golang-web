package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"web-crud-db/model"
	"web-crud-db/util"
)

func loadProperties() model.DbConfig {

	properties := util.ReadFile("resources/config.properties")

	return model.DbConfig{
		User:    properties["db-user"],
		Pass:    properties["db-passwd"],
		Net:     properties["db-net"],
		Address: properties["db-address"],
		DbName:  properties["db-schema"],
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
