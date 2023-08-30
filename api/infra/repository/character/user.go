package user

import (
	"fmt"

	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &Repository{db}
}

// 作成
func (r *Repository) Insert(user *entity.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// 条件付き取得
func (r *Repository) SelectById(user *entity.User, id string) error {
	if err := r.db.First(user, id).Error; err != nil {
		return err
	}

	return nil
}

// 条件付き更新
func (r *Repository) UpdateById(user *entity.User, id string) error {
	result := r.db.Model(user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// 条件付き削除
func (r *Repository) DeleteById(id string) error {
	result := r.db.Where("id= ? ", id).Delete(&entity.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
