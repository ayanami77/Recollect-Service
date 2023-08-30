package user

import "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"

type Repository interface {
	Insert(user *entity.User) error
	SelectById(user *entity.User, id string) error
	UpdateById(user *entity.User, id string) error
	DeleteById(id string) error
}
