package card

import "github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"

type Repository interface {
	Insert(card *entity.Card) error
	SelectById(card *entity.Card, id uint) error
	UpdateById(card *entity.Card, id uint) error
	DeleteById(id uint) error
}
