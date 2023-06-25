package adduser

import (
	"errors"

	"github.com/ThaksinCK/go-basic-api.git/config"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (r UserRepository) NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) AddUser(request AddUserRequest) (err error) {
	if config.BasicApiDatabase.DB == nil {
		return errors.New("database connection unavailable")
	}
	db := config.BasicApiDatabase.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("users").Create(&request).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}
