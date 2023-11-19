package userrepositoryimpl

import (
	"fmt"
	"github.com/karlosdaniel451/go-rest-api-template/domain/model"
	"github.com/karlosdaniel451/go-rest-api-template/errs"
	"gorm.io/gorm"
)

type UserRepositoryGORM struct {
	db *gorm.DB
}

func NewUserRepositoryGORM(db *gorm.DB) *UserRepositoryGORM {
	return &UserRepositoryGORM{db: db}
}

func (repository UserRepositoryGORM) Create(user *model.User) (*model.User, error) {
	result := repository.db.Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("it was not possible to insert user: %s", result.Error)
	}

	return user, nil
}

func (repository UserRepositoryGORM) GetById(id uint) (*model.User, error) {
	var user model.User

	result := repository.db.Preload("Tasks").First(&user, id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errs.NotFoundError{
				Message: fmt.Sprintf("there is no user with id %d", id),
			}
		}
		return nil, result.Error
	}

	return &user, nil
}

func (repository UserRepositoryGORM) GetByEmail(email string) (*model.User, error) {
	var user model.User

	result := repository.db.Preload("Tasks").First(&user, "email = ?", email)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errs.NotFoundError{
				Message: fmt.Sprintf("there is no user with email %s", email),
			}
		}
		return nil, result.Error
	}

	return &user, nil
}

func (repository UserRepositoryGORM) DeleteById(id uint) error {
	var user model.User

	result := repository.db.First(&user, id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return errs.NotFoundError{
				Message: fmt.Sprintf("there is no user with id %d", id),
			}
		}
		return result.Error
	}
	result = result.Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository UserRepositoryGORM) GetAll() ([]*model.User, error) {
	allUsers := make([]*model.User, 0)

	// result := repository.db.Find(&allUsers)
	result := repository.db.Model(&model.User{}).Preload("Tasks").Find(&allUsers)
	if result.Error != nil {
		return nil, result.Error
	}

	return allUsers, nil
}
