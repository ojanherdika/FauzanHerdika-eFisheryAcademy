package usecase

import (
	"e-commerce/entity"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type ICartUsecase interface {
	CreateCart(cart entity.CreateCartRequest) (entity.Cart, error)
	GetAllCart() ([]entity.Cart, error)
	GetCartById(id int) (entity.Cart, error)
	UpdateCart(cart entity.UpdateCartRequest, id int) (entity.Cart, error)
	DeleteCart(id int) error
}
type CartUsecase struct {
	cartRepository repository.ICartRespository
}

func NewCartUseCase(cartRepository repository.ICartRespository) *CartUsecase {
	return &CartUsecase{cartRepository}
}
func (usecase CartUsecase) CreateCart(cart entity.CreateCartRequest) (entity.CartResponse, error) {
	u := entity.Cart{
		Quantity:  cart.Quantity,
		Checkout:  cart.Checkout,
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
	}

	carts, err := usecase.cartRepository.Store(u)
	if err != nil {
		return entity.CartResponse{}, err
	}
	cartRes := entity.CartResponse{
		ID:        carts.ID,
		Quantity:  cart.Quantity,
		Checkout:  cart.Checkout,
		UserID:    cart.UserID,
		ProductID: cart.ProductID,
	}
	return cartRes, nil
}
func (usecase CartUsecase) GetAllCart() ([]entity.CartResponse, error) {
	carts, err := usecase.cartRepository.FindAll()
	if err != nil {
		return nil, err
	}
	cartRes := []entity.CartResponse{}

	//maping with copier
	copier.Copy(&cartRes, &carts)
	return cartRes, nil
}
func (usecase CartUsecase) GetCartById(id int) (entity.CartResponse, error) {
	cart, err := usecase.cartRepository.FindById(id)
	if err != nil {
		return entity.CartResponse{}, err
	}
	cartRes := entity.CartResponse{}
	copier.Copy(&cartRes, &cart)
	return cartRes, nil
}

func (usecase CartUsecase) UpdateCart(cartRequest entity.UpdateCartRequest, id int) (entity.CartResponse, error) {
	cart, err := usecase.cartRepository.FindById(id)
	if err != nil {
		return entity.CartResponse{}, err
	}

	// product.Name = productRequest.Name
	// product.Photo = productRequest.Photo
	// product.Price = productRequest.Price
	// product.Category = productRequest.Category
	// product.Description = productRequest.Description

	copier.CopyWithOption(&cart, &cartRequest, copier.Option{IgnoreEmpty: true})
	cart, err = usecase.cartRepository.Update(cart)
	cartRes := entity.CartResponse{}
	copier.Copy(&cartRes, &cart)
	return cartRes, nil
}
func (usecase CartUsecase) DeleteCart(id int) error {
	_, err := usecase.cartRepository.FindById(id)
	if err != nil {
		return err
	}
	err = usecase.cartRepository.Delete(id)
	return err
}
