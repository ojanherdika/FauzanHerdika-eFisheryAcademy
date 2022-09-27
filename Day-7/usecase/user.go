package usecase

import (
	"clean_arch/entity"
	"clean_arch/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	CreateUser(user entity.UserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
}
type UserUsecase struct {
	userRepository repository.IUserRespository
}

func NewUserUseCase(userRepository repository.IUserRespository) *UserUsecase {
	return &UserUsecase{userRepository}
}
func (usecase UserUsecase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	users, err := usecase.userRepository.Store(u)
	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		ID:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Phone:    users.Phone,
	}
	return userRes, nil
}
func (usecase UserUsecase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := usecase.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	// for _, value :=range []entity.UserResponse{
	// 	fmt.Println(value.ID)
	// }
	userRes := []entity.UserResponse{}
	// userRes = append(userRes, users)

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
