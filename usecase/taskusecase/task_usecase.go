package taskusecase

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
)

type TaskUseCase interface {
	Create(user *model.Task) (*model.Task, error)
	GetById(id uint) (*model.Task, error)
	GetByName(name string) ([]*model.Task, error)
	GetByDescription(description string) ([]*model.Task, error)
	DeleteById(id uint) error
	GetAll() ([]*model.Task, error)
}
