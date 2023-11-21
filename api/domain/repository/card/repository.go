package card

import "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Repository interface {
	Insert(card *entity.Card) error
	BatchInsert(cards *[]entity.Card) error
	SelectAll(card *[]entity.Card, userID string) error
	UpdateById(card *entity.Card, id string) error
	DeleteById(id string) error
}
