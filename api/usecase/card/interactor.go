package card

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	cardRepository "github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
)

//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Interactor interface {
	ListCards(sub string) ([]entity.Card, error)
	CreateCard(card entity.Card) (entity.Card, error)
	CreateCards(cards []entity.Card) ([]entity.Card, error)
	UpdateCard(card entity.Card, id string, sub string) (entity.Card, error)
	DeleteCard(id string, sub string) error
}

type interactor struct {
	cardRepository cardRepository.Repository
}

func New(cardRepository cardRepository.Repository) Interactor {
	return &interactor{cardRepository}
}

func (i *interactor) ListCards(sub string) ([]entity.Card, error) {
	cards := []entity.Card{}

	err := i.cardRepository.SelectAll(&cards, sub)
	if err != nil {
		return []entity.Card{}, myerror.InternalServerError
	}

	return cards, nil
}

func (i *interactor) CreateCard(card entity.Card) (entity.Card, error) {
	err := i.cardRepository.Insert(&card)
	if err != nil {
		return entity.Card{}, myerror.InternalServerError
	}

	return card, nil
}

func (i *interactor) CreateCards(cards []entity.Card) ([]entity.Card, error) {
	err := i.cardRepository.BatchInsert(&cards)
	if err != nil {
		return []entity.Card{}, myerror.InternalServerError
	}

	return cards, nil
}

func (i *interactor) UpdateCard(card entity.Card, id string, sub string) (entity.Card, error) {
	if err := i.cardRepository.UpdateById(&card, id, sub); err != nil {
		return entity.Card{}, myerror.InternalServerError
	}

	return card, nil
}

func (i *interactor) DeleteCard(id string, sub string) error {
	if err := i.cardRepository.DeleteById(id, sub); err != nil {
		return myerror.InternalServerError
	}

	return nil
}
