package userusecaseimpl

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/repository/userrepository"
)

type UserUseCaseImpl struct {
	repository userrepository.UserRepository
}

func NewUserUseCaseImpl(userRepository userrepository.UserRepository) UserUseCaseImpl {
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
