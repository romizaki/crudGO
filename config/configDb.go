package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDbConnection is creating connection
func SetupDbConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Env tidak connect")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	// dsn := "romizaki:romizaki@tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to create a connection to database")
	}
	// db.AutoMigrate(&entity.Item{}, &entity.Order{})
	return db
}

//CloseDbConnection close connection
func CloseDbConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection database")
	}
	dbSQL.Close()
}
