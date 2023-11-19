package usecase

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	userRepository "github.com/karlosdaniel451/go-rest-api-template/repository"
)

type UserUseCase interface {
	Create(user *model.User) (*model.User, error)
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	DeleteById(id uint) error
	GetAll() ([]*model.User, error)
}

type UserUseCaseImpl struct {
	repository userRepository.UserRepository
}

func NewUserUseCaseImpl(userRepository userRepository.UserRepository) UserUseCaseImpl {
	return UserUseCaseImpl{repository: userRepository}
}

func (useCase UserUseCaseImpl) Create(user *model.User) (*model.User, error) {
	return useCase.repository.Create(user)
}

func (useCase UserUseCaseImpl) GetById(id uint) (*model.User, error) {
	return useCase.repository.GetById(id)
}

func (useCase UserUseCaseImpl) GetByEmail(email string) (*model.User, error) {
	return useCase.repository.GetByEmail(email)
}

func (useCase UserUseCaseImpl) DeleteById(id uint) error {
	return useCase.repository.DeleteById(id)
}

func (useCase UserUseCaseImpl) GetAll() ([]*model.User, error) {
	return useCase.repository.GetAll()
}
