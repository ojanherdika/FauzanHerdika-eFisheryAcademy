package repository

import (
	"clean_arch/entity"

	"gorm.io/gorm"
)

type IUserRespository interface {
	Store(user entity.User) (entity.User, error)
	FindAll() ([]entity.User, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
func (r UserRepository) Store(user entity.User) (entity.User, error) {
	if err := r.db.Debug().Create(&user).Error; err != nil {

	}
	return user, nil
}
func (r UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Debug().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
