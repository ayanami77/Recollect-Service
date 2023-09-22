package card

import (
	"fmt"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/card"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) card.Repository {
	return &Repository{db}
}

func (r *Repository) Insert(card *entity.Card) error {
	if err := r.db.Create(card).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectAll(cards *[]entity.Card, userID string) error {
	if err := r.db.Where("user_id = ?", userID).Find(cards).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateById(card *entity.Card, id string) error {
	result := r.db.Model(card).Where("card_id = ?", id).Updates(card)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (r *Repository) DeleteById(id string) error {
	result := r.db.Where("card_id = ?", id).Delete(&entity.Card{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
