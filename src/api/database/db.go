package database

import (
	"config"
	"github.com/jinzhu/gorm"

)
// Connect connect to databse
func Connect() (*gorm.DB, error)  {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}