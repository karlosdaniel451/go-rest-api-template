package taskrepository

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
)

type TaskRepository interface {
	Create(task *model.Task) (*model.Task, error)
	GetById(id uint) (*model.Task, error)
	GetByName(name string) ([]*model.Task, error)
	GetByDescription(description string) ([]*model.Task, error)
	DeleteById(id uint) error
	GetAll() ([]*model.Task, error)
}
