package user

import "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Repository interface {
	Insert(user *entity.User) error
	SelectById(user *entity.User, id string) error
	UpdateById(user *entity.User, id string) error
	DeleteById(id string) error
}
