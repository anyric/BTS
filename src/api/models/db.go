package models

import (
	"github.com/anyric/bts/src/config"
	"github.com/jinzhu/gorm"
)

// Connect to a databse an return an instance
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}
