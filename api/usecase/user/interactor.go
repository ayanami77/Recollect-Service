package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	userRepository "github.com/Seiya-Tagami/Recollect-Service/api/domain/repository/user"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/util/openaiutil"
)

//go:generate mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Interactor interface {
	GetUser(sub string) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User, sub string) (entity.User, error)
	DeleteUser(sub string) error
	CheckEmailDuplication(email string) (bool, error)
	CheckUserIDDuplication(userID string) (bool, error)
	AnalyzeUserHistory(sub string) (entity.User, error)
}

type interactor struct {
	userRepository userRepository.Repository
}

func New(userRepository userRepository.Repository) Interactor {
	return &interactor{userRepository}
}

func (i *interactor) GetUser(sub string) (entity.User, error) {
	user := entity.User{}

	err := i.userRepository.SelectBySub(&user, sub)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	return user, nil
}

func (i *interactor) CreateUser(user entity.User) (entity.User, error) {
	err := i.userRepository.Insert(&user)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	return user, nil
}

func (i *interactor) UpdateUser(user entity.User, sub string) (entity.User, error) {
	if err := i.userRepository.UpdateBySub(&user, sub); err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	return user, nil
}

func (i *interactor) DeleteUser(sub string) error {
	if err := i.userRepository.DeleteBySub(sub); err != nil {
		return myerror.InternalServerError
	}

	return nil
}

func (i *interactor) CheckEmailDuplication(email string) (bool, error) {
	isDuplicated, err := i.userRepository.ExistsByEmail(email)
	if err != nil {
		return false, myerror.InternalServerError
	}

	return isDuplicated, nil
}

func (i *interactor) CheckUserIDDuplication(userID string) (bool, error) {
	isDuplicated, err := i.userRepository.ExistsByUserID(userID)
	if err != nil {
		return false, myerror.InternalServerError
	}

	return isDuplicated, nil
}

func (i *interactor) AnalyzeUserHistory(sub string) (entity.User, error) {
	analysisData, err := i.userRepository.GetAnalysisResultStringBySub(sub)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	comprehensiveAnalysisResult, err := getComprehensiveAnalysisResult(analysisData)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	comprehensiveAnalysisScore, err := getComprehensiveAnalysisScore(comprehensiveAnalysisResult, analysisData)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	user := entity.User{}

	user.ComprehensiveAnalysisResult = comprehensiveAnalysisResult
	user.ComprehensiveAnalysisScore = comprehensiveAnalysisScore

	return i.UpdateUser(user, sub)
}

func getComprehensiveAnalysisResult(analysisData userRepository.AnalysisData) (string, error) {
	prompt := `
    下記の特性と自分史から、その人を分析しフォーマット例のように一言にまとめ、マークダウンで出力してください。
    
    フォーマット例:「
        **実験とリーダーシップの熱心な探求者**\n実験好きの特性は、新しいことへの挑戦と知識の追求を示しており、リーダーシップの特性は、チームを導き、目標達成に向けて取り組む力を表しています。また、計画性、積極性、コミュニケーション能力もこのフレーズに含まれており、あなたの多面的な資質を総合的に表現しています。
    」

    特性:「
	` + analysisData.AnalysisResultString + `
    」

	自分史: 「
    ` + analysisData.UserHistoryString + `
    」
	`
	return openaiutil.FetchOpenAIResponse(prompt)
}

func getComprehensiveAnalysisScore(comprehensiveAnalysisResult string, analysisData userRepository.AnalysisData) (string, error) {
	prompt := `
     下記の分析結果は、その人を総合分析したものです。重要なので忘れないでください。
    分析結果:「
	` + comprehensiveAnalysisResult + `
    」

    また、その人の特性と自分史は以下の通りです。
    特性:「
	` + analysisData.AnalysisResultString + `
    」

	自分史: 「
    ` + analysisData.UserHistoryString + `
    」

    これらの情報から、特性を抽出または新しく作成し、50～100点で点数化します。
    特に点数が高いものから、必ず6つ表示してください。
    フォーマット例:「
    - **責任感**: __50__\n- **リーダーシップ**: __70__\n- **努力家**: __95__\n- **実験好き**: __84__\n- **チームワーク**: __72__\n- **コミュニケーション能力**: __92__
    」
  	`
	return openaiutil.FetchOpenAIResponse(prompt)
}
