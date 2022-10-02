package usecase

import (
	"e-commerce/entity"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	CreateUser(user entity.CreateUserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateUser(user entity.UpdateUserRequset, id int) (entity.User, error)
	DeleteUser(id int) error
}
type UserUsecase struct {
	userRepository repository.IUserRespository
}

func NewUserUseCase(userRepository repository.IUserRespository) *UserUsecase {
	return &UserUsecase{userRepository}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func (usecase UserUsecase) CreateUser(user entity.CreateUserRequest) (entity.UserResponse, error) {
	hash, err := HashPassword(user.Password)
	u := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hash,
		Phone:    user.Phone,
	}

	users, err := usecase.userRepository.Store(u)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		ID:    users.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}
	return userRes, nil
}
func (usecase UserUsecase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := usecase.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	userRes := []entity.UserResponse{}

	//maping with copier
	copier.Copy(&userRes, &users)
	return userRes, nil
}
func (usecase UserUsecase) GetUserById(id int) (entity.UserResponse, error) {
	user, err := usecase.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, nil
}

func (usecase UserUsecase) UpdateUser(userRequest entity.UpdateUserRequset, id int) (entity.UserResponse, error) {
	user, err := usecase.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}
	hash, err := HashPassword(userRequest.Password)

	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Password = hash
	user.Phone = userRequest.Phone

	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true})
	user, err = usecase.userRepository.Update(user)
	userRes := entity.UserResponse{}
	copier.Copy(&userRes, &user)
	return userRes, nil
}
func (usecase UserUsecase) DeleteUser(id int) error {
	_, err := usecase.userRepository.FindById(id)
	if err != nil {
		return err
	}
	err = usecase.userRepository.Delete(id)
	return err
}
