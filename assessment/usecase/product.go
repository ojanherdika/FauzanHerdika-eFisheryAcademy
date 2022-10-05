package usecase

import (
	"e-commerce/entity"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type IProductUsecase interface {
	CreateProduct(product entity.CreateProductRequest) (entity.Product, error)
	GetAllProduct() ([]entity.Product, error)
	GetProductById(id int) (entity.Product, error)
	GetProductByCategory(category string) ([]entity.Product, error)
	GetProductByPrice(priceMin int, priceMax int) ([]entity.Product, error)
	UpdateProduct(product entity.UpdateProductRequest, id int) (entity.Product, error)
	DeleteProduct(id int) error
}
type ProductUsecase struct {
	productRepository repository.IProductRespository
}

func NewProductUseCase(productRepository repository.IProductRespository) *ProductUsecase {
	return &ProductUsecase{productRepository}
}
func (usecase ProductUsecase) CreateProduct(product entity.CreateProductRequest) (entity.ProductResponse, error) {
	u := entity.Product{
		Name:        product.Name,
		Photo:       product.Photo,
		Price:       product.Price,
		Category:    product.Category,
		Description: product.Description,
	}

	products, err := usecase.productRepository.Store(u)
	if err != nil {
		return entity.ProductResponse{}, err
	}
	productRes := entity.ProductResponse{
		ID:          products.ID,
		Name:        products.Name,
		Photo:       products.Photo,
		Price:       products.Price,
		Category:    products.Category,
		Description: products.Description,
	}
	return productRes, nil
}
func (usecase ProductUsecase) GetAllProduct() ([]entity.ProductResponse, error) {
	products, err := usecase.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	productRes := []entity.ProductResponse{}

	//maping with copier
	copier.Copy(&productRes, &products)
	return productRes, nil
}
func (usecase ProductUsecase) GetProductById(id int) (entity.ProductResponse, error) {
	product, err := usecase.productRepository.FindById(id)
	if err != nil {
		return entity.ProductResponse{}, err
	}
	productRes := entity.ProductResponse{}
	copier.Copy(&productRes, &product)
	return productRes, nil
}
func (usecase ProductUsecase) GetProductByCategory(category string) ([]entity.ProductResponse, error) {
	product, err := usecase.productRepository.FindByCategory(category)
	if err != nil {
		return []entity.ProductResponse{}, err
	}
	productRes := []entity.ProductResponse{}
	copier.Copy(&productRes, &product)
	return productRes, nil
}
func (usecase ProductUsecase) GetProductByPrice(priceMin int, priceMax int) ([]entity.ProductResponse, error) {
	product, err := usecase.productRepository.FindByPrice(priceMin, priceMax)
	if err != nil {
		return []entity.ProductResponse{}, err
	}
	productRes := []entity.ProductResponse{}
	copier.Copy(&productRes, &product)
	return productRes, nil
}
func (usecase ProductUsecase) UpdateProduct(productRequest entity.UpdateProductRequest, id int) (entity.ProductResponse, error) {
	product, err := usecase.productRepository.FindById(id)
	if err != nil {
		return entity.ProductResponse{}, err
	}

	product.Name = productRequest.Name
	product.Photo = productRequest.Photo
	product.Price = productRequest.Price
	product.Category = productRequest.Category
	product.Description = productRequest.Description

	copier.CopyWithOption(&product, &productRequest, copier.Option{IgnoreEmpty: true})
	product, err = usecase.productRepository.Update(product)
	productRes := entity.ProductResponse{}
	copier.Copy(&productRes, &product)
	return productRes, nil
}
func (usecase ProductUsecase) DeleteProduct(id int) error {
	_, err := usecase.productRepository.FindById(id)
	if err != nil {
		return err
	}
	err = usecase.productRepository.Delete(id)
	return err
}
