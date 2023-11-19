package userusecase

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
)

type UserUseCase interface {
	Create(user *model.User) (*model.User, error)
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	DeleteById(id uint) error
	GetAll() ([]*model.User, error)
}
