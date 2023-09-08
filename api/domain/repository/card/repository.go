package card

import "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"

type Repository interface {
	Insert(card *entity.Card) error
	SelectAll(card *[]entity.Card) error
	SelectById(card *entity.Card, id string) error
	UpdateById(card *entity.Card, id string) error
	DeleteById(id string) error
}
