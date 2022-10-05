package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type ICartRespository interface {
	Store(cart entity.Cart) (entity.Cart, error)
	FindAll() ([]entity.Cart, error)
	FindByUser(id int) ([]entity.Cart, error)
	FindById(id int) (entity.Cart, error)
	Update(cart entity.Cart) (entity.Cart, error)
	Delete(id int) error
}
type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db}
}
func (r CartRepository) Store(cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Debug().Create(&cart).Error; err != nil {

	}
	return cart, nil
}
func (r CartRepository) FindAll() ([]entity.Cart, error) {
	var carts []entity.Cart
	if err := r.db.Debug().Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}
func (r CartRepository) FindByUser(id int) ([]entity.Cart, error) {
	var cart []entity.Cart
	if err := r.db.Debug().Where("user_id = ?AND checkout = ?", id, false).Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
func (r CartRepository) FindById(id int) (entity.Cart, error) {
	var cart entity.Cart
	if err := r.db.Debug().Where("id = ?", id).First(&cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}
func (repo CartRepository) Update(cart entity.Cart) (entity.Cart, error) {
	if err := repo.db.Debug().Save(&cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (repo CartRepository) Delete(id int) error {
	if err := repo.db.Debug().Delete(&entity.Cart{}, id).Error; err != nil {
		return err
	}
	return nil
}
