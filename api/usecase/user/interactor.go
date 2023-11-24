package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	userRepository "github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/user"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Interactor interface {
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User, sub string) (entity.User, error)
	DeleteUser(sub string) error
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

func (i *interactor) UpdateUser(user entity.User, sub string) (entity.User, error) {
	if err := i.userRepository.UpdateBySub(&user, sub); err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	return user, nil
}

func (i *interactor) DeleteUser(sub string) error {
	if err := i.userRepository.DeleteBySub(sub); err != nil {
		return myerror.InternalServerError
	}

	return nil
}

func (i *interactor) CheckEmailDuplication(email string) (bool, error) {
	isDuplicated, err := i.userRepository.ExistsByEmail(email)
	if err != nil {
		return false, myerror.InternalServerError
	}

	return isDuplicated, nil
}

func (i *interactor) CheckUserIDDuplication(userID string) (bool, error) {
	isDuplicated, err := i.userRepository.ExistsByUserID(userID)
	if err != nil {
		return false, myerror.InternalServerError
	}

	return isDuplicated, nil
}
