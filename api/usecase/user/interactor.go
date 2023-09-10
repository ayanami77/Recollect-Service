package user

import (
	"fmt"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	userRepository "github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/user"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Interactor interface {
	GetUser(id string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User, id string) (entity.User, error)
	DeleteUser(id string) error
	LoginUser(id string, password string) (string, error)
	//LogoutUser(id string) (entity.User, error)
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

func (i *interactor) LoginUser(id string, password string) (string, error) {
	user := entity.User{}

	err := i.userRepository.SelectById(&user, id)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		err = fmt.Errorf("password is not correct")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//func (i *interactor) LogoutUser(id string) (entity.User, error) {}
