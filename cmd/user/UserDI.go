package user

import "gorm.io/gorm"

func ConfigureDependencies(db *gorm.DB) UserController {
	userRepo := UserRepository{
		Db: db,
	}
	userSvc := UserService{
		UserRepository: userRepo,
	}
	userCtrl := UserController{
		UserService: userSvc,
	}

	return userCtrl
}
