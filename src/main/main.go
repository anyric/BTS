package main

import (
	"fmt"
	"log"
	"api"
	"api/models"
	"config"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db := ConnectToDatabase()
	defer db.Close()
    api.Run()
}

// ConnectToDatabase opens database connection
func ConnectToDatabase() *gorm.DB {
	e := godotenv.Load()
	config.Load()
	if e != nil {
		log.Fatalln(fmt.Errorf("Unable to load .env: %v", e))
	}
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		fmt.Print(err)
	}
	db.Debug().AutoMigrate(&models.Account{})
	models.SetDatabase(db)
	return db
}