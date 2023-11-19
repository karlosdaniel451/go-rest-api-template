package taskusecaseimpl

import (
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/repository/taskrepository"
)

type TaskUseCaseImpl struct {
	taskRepository taskrepository.TaskRepository
}

func NewTaskUseCaseImpl(taskRepository taskrepository.TaskRepository) TaskUseCaseImpl {
	return TaskUseCaseImpl{taskRepository: taskRepository}
}

func (useCase TaskUseCaseImpl) Create(user *model.Task) (*model.Task, error) {
	return useCase.taskRepository.Create(user)
}

func (useCase TaskUseCaseImpl) GetById(id uint) (*model.Task, error) {
	return useCase.taskRepository.GetById(id)
}

func (useCase TaskUseCaseImpl) GetByName(name string) ([]*model.Task, error) {
	return useCase.taskRepository.GetByName(name)
}

func (useCase TaskUseCaseImpl) GetByDescription(description string) ([]*model.Task, error) {
	return useCase.taskRepository.GetByDescription(description)
}

func (useCase TaskUseCaseImpl) DeleteById(id uint) error {
	return useCase.taskRepository.DeleteById(id)
}

func (useCase TaskUseCaseImpl) GetAll() ([]*model.Task, error) {
	return useCase.taskRepository.GetAll()
}
