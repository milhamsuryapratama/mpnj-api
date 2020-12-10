package config

import (
	"github.com/joho/godotenv"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ..
func Connect() *gorm.DB {
	godotenv.Load("./.env")
	dbName := os.Getenv("MYSQL_DB")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db
}
