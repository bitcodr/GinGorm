package config

import (
	"fmt"
	"log"
	"os"
	 _ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

//InitDB database
func InitDB() {


	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, pass, host, port, database)

		log.Println(dbInfo)
	db, err = gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("error in database connection")
	}
}

//GetDB func
func GetDB() *gorm.DB {
	return db
}

