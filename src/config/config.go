package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// declare port variable
var (
	PORT = 0
	DBDRIVER = ""
	DBURL = ""
)

// Load the variables from .env file
func Load() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		log.Println(err)
		PORT = 9000
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
}
