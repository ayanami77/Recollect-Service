package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	userRepository "github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/user"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
)

type Interactor interface {
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User, id string) (entity.User, error)
	DeleteUser(id string) error
	CheckEmailDuplication(email string) (bool, error)
	CheckUserIDDuplication(userID string) (bool, error)
}

type interactor struct {
	userRepository userRepository.Repository
}

func New(userRepository userRepository.Repository) Interactor {
	return &interactor{userRepository}
}

func (i *interactor) CreateUser(user entity.User) (entity.User, error) {
	err := i.userRepository.Insert(&user)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	return user, nil
}

func (i *interactor) UpdateUser(user entity.User, id string) (entity.User, error) {
	if err := i.userRepository.UpdateById(&user, id); err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	return user, nil
}

func (i *interactor) DeleteUser(id string) error {
	if err := i.userRepository.DeleteById(id); err != nil {
		return myerror.InternalServerError
	}

	return nil
}

func (i *interactor) CheckEmailDuplication(email string) (bool, error) {
	user := entity.User{}
	if err := i.userRepository.SelectByEmail(&user, email); err != nil {
		return false, myerror.InternalServerError
	}

	return user.Email != email, nil
}

func (i *interactor) CheckUserIDDuplication(userID string) (bool, error) {
	user := entity.User{}
	if err := i.userRepository.SelectByUserID(&user, userID); err != nil {
		return false, myerror.InternalServerError
	}

	return user.UserID != userID, nil
}
