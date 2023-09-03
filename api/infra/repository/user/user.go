package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"

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

func (r *Repository) Insert(user *entity.User) error {
	if len(user.UserName) == 0 {
		user.UserName = user.UserID
	}

	validate := validator.New()
	//validate.RegisterValidation("includeNumeric", entity.IncludeAlphabetic)
	//validate.RegisterValidation("includeAlphabetic", entity.IncludeNumeric)
	if err := validate.Struct(user); err != nil {
		return err
	}

	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectById(user *entity.User, id string) error {
	if err := r.db.First(user, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateById(user *entity.User, id string) error {
	validate := validator.New()
	//validate.RegisterValidation("includeNumeric", entity.IncludeAlphabetic)
	//validate.RegisterValidation("includeAlphabetic", entity.IncludeNumeric)
	if err := validate.Struct(user); err != nil {
		return err
	}

	result := r.db.Model(user).Where("user_id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (r *Repository) DeleteById(id string) error {
	result := r.db.Where("user_id = ? ", id).Delete(&entity.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
