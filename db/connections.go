package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB        *gorm.DB
	dbUrlProd string
	dbUrlDev  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// Extracting MYSQL Database
	dbUrlProd = "root:" + os.Getenv("MYSQL_ROOT_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DATABASE")
	dbUrlDev = os.Getenv("DB_URL")
	if os.Getenv("MODE") == "development" {
		DB, err = gorm.Open(mysql.Open(dbUrlDev), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	} else {
		DB, err = gorm.Open(mysql.Open(dbUrlProd), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}
}
