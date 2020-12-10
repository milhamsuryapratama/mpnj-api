package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ..
func Connect() *gorm.DB {
	//godotenv.Load("./.env")
	dbName := "mpnj"
	dbUser := "root"
	dbPassword := ""
	dsn := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db
}
