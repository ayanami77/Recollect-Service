package user

import (
	"fmt"
	userUsecase "github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
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

func (r *Repository) SelectBySub(user *entity.User, sub string) error {
	result := r.db.Where("sub = ?", sub).First(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *Repository) UpdateBySub(user *entity.User, sub string) error {
	//TODO: user_idなど、本来validateが必要なものを更新するusecaseがある時は、必ずONに戻す。今は総合分析のみなので許容
	//validate := validator.New()
	//validate.RegisterValidation("includeNumeric", entity.IncludeAlphabetic)
	//validate.RegisterValidation("includeAlphabetic", entity.IncludeNumeric)
	//if err := validate.Struct(user); err != nil {
	//	return err
	//}

	result := r.db.Model(user).Where("sub = ?", sub).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (r *Repository) DeleteBySub(sub string) error {
	result := r.db.Where("sub = ? ", sub).Delete(&entity.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (r *Repository) ExistsByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) ExistsByUserID(userID string) (bool, error) {
	var count int64
	if err := r.db.Model(&entity.User{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) GetAnalysisDataBySub(sub string) (userUsecase.AnalysisData, error) {
	cards := []entity.Card{}
	if err := r.db.Select("AnalysisResult").Find(&cards).Where("sub = ?", sub).Error; err != nil {
		return userUsecase.AnalysisData{}, err
	}

	userHistoryString := ""
	analysisResultString := ""
	for _, card := range cards {
		analysisResultString += card.AnalysisResult + "\n"
		userHistoryString += card.Title + "\n" + card.Content + "\n"
	}

	analysisString := userUsecase.AnalysisData{
		UserHistoryString:    userHistoryString,
		AnalysisResultString: analysisResultString,
	}

	return analysisString, nil
}
