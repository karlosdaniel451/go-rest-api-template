package taskrepositoryimpl

import (
	"fmt"
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/errs"
	"gorm.io/gorm"
)

type TaskRepositoryGORM struct {
	db *gorm.DB
}

func NewTaskRepositoryGORM(db *gorm.DB) *TaskRepositoryGORM {
	return &TaskRepositoryGORM{db: db}
}

func (repository TaskRepositoryGORM) Create(task *model.Task) (*model.Task, error) {
	result := repository.db.Create(task)
	if result.Error != nil {
		return nil, fmt.Errorf("it was not possible to insert task: %s", result.Error)
	}

	return task, nil
}

func (repository TaskRepositoryGORM) GetById(id uint) (*model.Task, error) {
	var task model.Task

	result := repository.db.First(&task, "id = ?", id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errs.NotFoundError{
				Message: fmt.Sprintf("there is no task with id %d", id),
			}
		}
		return nil, result.Error
	}

	return &task, nil
}

func (repository TaskRepositoryGORM) GetByName(name string) ([]*model.Task, error) {
	tasks := make([]*model.Task, 0)

	result := repository.db.Where("NAME LIKE %?%", name)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (repository TaskRepositoryGORM) GetByDescription(description string) ([]*model.Task, error) {
	tasks := make([]*model.Task, 0)

	result := repository.db.Where("description LIKE %?%", description)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (repository TaskRepositoryGORM) DeleteById(id uint) error {
	var task model.Task

	result := repository.db.First(&task, id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return errs.NotFoundError{
				Message: fmt.Sprintf("there is no task with id %d", id),
			}
		}
		return result.Error
	}
	result = result.Delete(&task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository TaskRepositoryGORM) GetAll() ([]*model.Task, error) {
	allTasks := make([]*model.Task, 0)

	result := repository.db.Find(&allTasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return allTasks, nil
}
