package user

import (
	"fmt"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/user"
	"github.com/go-playground/validator/v10"
	"github.com/pkoukk/tiktoken-go"
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

func (r *Repository) GetAnalysisDataBySub(sub string) (user.AnalysisData, error) {
	cards := []entity.Card{}
	if err := r.db.Select("Title", "AnalysisResult").Find(&cards).Where("sub = ?", sub).Error; err != nil {
		return user.AnalysisData{}, err
	}

	availableTokens := 3700
	cardTitleString := ""
	analysisResultString := ""

	// token数取得の準備
	encoding := "cl100k_base"
	tkm, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return user.AnalysisData{}, err
	}

	for _, card := range cards {
		if len(tkm.Encode(cardTitleString, nil, nil)) >= 500 {
			break
		}
		cardTitleString += card.Title + "\n"
		fmt.Println(card)
	}
	fmt.Println(cardTitleString)
	fmt.Printf("cardTitleString Token: %v\n", len(tkm.Encode(cardTitleString, nil, nil)))
	availableTokens = availableTokens - len(tkm.Encode(cardTitleString, nil, nil))

	for _, card := range cards {
		// AnalysisResultを文字列に追加すると、利用可能トークンを超える場合、追加しない
		if availableTokens-len(tkm.Encode(analysisResultString+card.AnalysisResult+"\n", nil, nil)) < 0 {
			break
		}
		analysisResultString += card.AnalysisResult + "\n"
		availableTokens -= len(tkm.Encode(analysisResultString, nil, nil))
		fmt.Printf("analysisResultString Token: %v\n", len(tkm.Encode(analysisResultString, nil, nil)))
	}

	fmt.Println(availableTokens)

	analysisString := user.AnalysisData{
		CardTitleString:      cardTitleString,
		AnalysisResultString: analysisResultString,
	}

	return analysisString, nil
}
