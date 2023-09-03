package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	userRepository "github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/user"
)

type Interactor interface {
	GetUser(id string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User, id string) (entity.User, error)
	DeleteUser(id string) error
}

type interactor struct {
	userRepository userRepository.Repository
}

func New(userRepository userRepository.Repository) Interactor {
	return &interactor{userRepository}
}

func (i *interactor) GetUser(id string) (entity.User, error) {
	user := entity.User{}

	err := i.userRepository.SelectById(&user, id)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (i *interactor) CreateUser(user entity.User) (entity.User, error) {
	err := i.userRepository.Insert(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (i *interactor) UpdateUser(user entity.User, id string) (entity.User, error) {
	if err := i.userRepository.UpdateById(&user, id); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (i *interactor) DeleteUser(id string) error {
	if err := i.userRepository.DeleteById(id); err != nil {
		return err
	}

	return nil
}
