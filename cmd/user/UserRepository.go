package user

import (
	"github.com/guilleamutio/go4money/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		Db: db,
	}
}

func (userRepo *UserRepository) createUser(user User) {
	userRepo.Db.Create(&database.User{Username: user.Username, Password: user.Password})
}
