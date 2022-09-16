package user

import (
	"gorm.io/gorm"
)

func InitializeDependencies(db *gorm.DB) UserController {
	userRepo := NewUserRepository(db)
	userSvc := NewUserService(userRepo)
	userCtrl := NewUserController(userSvc)
	return userCtrl
}
