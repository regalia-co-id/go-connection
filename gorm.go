package connection

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func MySQL() *gorm.DB {
	err := godotenv.Load(".env.local")
	if err != nil {
		godotenv.Load(".env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	db, err := gorm.Open("mysql", username+":"+password+"@("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")

	db.DB().SetConnMaxLifetime(time.Minute * 10)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(0)

	if err != nil {
		log.Fatal("error:", err)
	}

	return db
}
