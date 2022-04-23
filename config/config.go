package config

import (
	"app-golang/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := model.Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8888"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "admin"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "admin12345"),
		DB_NAME:        GetOrDefault("DB_NAME", "training"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		DB_HOST:        GetOrDefault("DB_HOST", "mysqlgolang.cv8cnjvlnjwz.us-west-1.rds.amazonaws.com"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Book{})
}

func GetOrDefault(envKey, defaultValue string) string {
	if val, exist := os.LookupEnv(envKey); exist {
		return val
	}

	return defaultValue
}
