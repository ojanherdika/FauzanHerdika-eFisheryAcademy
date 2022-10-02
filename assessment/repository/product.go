package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type IProductRespository interface {
	Store(product entity.Product) (entity.Product, error)
	FindAll() ([]entity.Product, error)
	FindById(id int) (entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	Delete(id int) error
}
type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}
func (r ProductRepository) Store(product entity.Product) (entity.Product, error) {
	if err := r.db.Debug().Create(&product).Error; err != nil {

	}
	return product, nil
}
func (r ProductRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	if err := r.db.Debug().Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (r ProductRepository) FindById(id int) (entity.Product, error) {
	var user entity.Product
	if err := r.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return entity.Product{}, err
	}
	return user, nil
}
func (repo ProductRepository) Update(product entity.Product) (entity.Product, error) {
	if err := repo.db.Debug().Save(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (repo ProductRepository) Delete(id int) error {
	if err := repo.db.Debug().Delete(&entity.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
