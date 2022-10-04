package usecase

import (
	"e-commerce/config"
	"e-commerce/entity"
	"e-commerce/repository"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
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

func getUserByID(e int) (entity.UserResponse, error) {
	var user entity.UserResponse

	if err := config.DB.Model(entity.User{}).Where("id = ?", e).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
func getProductByID(e int) (entity.ProductResponse, error) {
	var product entity.ProductResponse

	if err := config.DB.Model(entity.Product{}).Where("id = ?", e).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		}
		return product, err
	}
	return product, nil
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
	user, _ := getUserByID(cart.UserID)
	product, _ := getProductByID(cart.ProductID)
	// convertedUser, _ := json.Marshal(user)
	cartRes := entity.CartResponse{
		ID:        carts.ID,
		Quantity:  cart.Quantity,
		Checkout:  cart.Checkout,
		UserID:    cart.UserID,
		User:      user,
		ProductID: cart.ProductID,
		Product:   product,
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
