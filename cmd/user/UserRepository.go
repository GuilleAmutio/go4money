package user

import (
	"github.com/guilleamutio/go4money/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (userRep *UserRepository) HastaLaPolla(user User) {
	userRep.Db.Create(&database.User{Username: user.Username, Password: user.Password})
}
