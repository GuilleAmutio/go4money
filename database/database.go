package database

import (
	"github.com/guilleamutio/go4money/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func OpenDatabase(config util.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Println(err)
	}

	return db, err

}
