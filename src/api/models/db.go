package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetDatabase  set the open database connection
func SetDatabase(database *gorm.DB) {
	db = database
}
