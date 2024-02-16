package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

var db *sql.DB

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	return viper.GetString(key)
}

func ConnectDB() {
	var err error
	dbHost := GetConfig("DB_HOST")
	dbPort := GetConfig("DB_PORT")
	dbUser := GetConfig("DB_USERNAME")
	dbPass := GetConfig("DB_PASSWORD")
	dbName := GetConfig("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))
	if err != nil {
		fmt.Println("Error connecting to database", db)
		panic(err.Error())
	}
}

func GetDB() *sql.DB {
	return db
}
