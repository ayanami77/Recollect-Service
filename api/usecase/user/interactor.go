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
	analysisResultString, err := i.userRepository.GetAnalysisResultStringBySub(sub)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	comprehensiveAnalysisResult, err := getComprehensiveAnalysisResult(analysisResultString)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	comprehensiveAnalysisScore, err := getComprehensiveAnalysisScore(comprehensiveAnalysisResult)
	if err != nil {
		return entity.User{}, myerror.InternalServerError
	}

	user := entity.User{}

	user.ComprehensiveAnalysisResult = comprehensiveAnalysisResult
	user.ComprehensiveAnalysisScore = comprehensiveAnalysisScore

	return i.UpdateUser(user, sub)
}

func getComprehensiveAnalysisResult(analysisResultString string) (string, error) {
	prompt := `
    下記の特性から、その人を分析しフォーマット例のように一言にまとめ、マークダウンで出力してください。
    
    フォーマット例:「
        **実験とリーダーシップの熱心な探求者**\n実験好きの特性は、新しいことへの挑戦と知識の追求を示しており、リーダーシップの特性は、チームを導き、目標達成に向けて取り組む力を表しています。また、計画性、積極性、コミュニケーション能力もこのフレーズに含まれており、あなたの多面的な資質を総合的に表現しています。
    」

    特性:「
	` + analysisResultString + `
    」
  	`
	return openaiutil.FetchOpenAIResponse(prompt)
}

func getComprehensiveAnalysisScore(comprehensiveAnalysisResult string) (string, error) {
	prompt := `
     下記の分析結果は、その人を総合分析したものです。重要なので忘れないでください。
    分析結果:「
	` + comprehensiveAnalysisResult + `
    」
    また、その人の特性一覧は以下の通りです。
    特性:「
    - **実験好き**: 文章の中で毎週実験をしていたことや実験を通じてわくわく感を感じていたことから、実験に対する興味や好奇心があることが分かります。\n- **チームワーク**: 化学実験では部員と協力し、文化祭の準備期間でも自分たちで何をするか考えて取り組んでいたことから、チームでの協力や協調性を大切にする特性が見受けられます。\n- **計画的**: 文化祭の準備期間ではじっくりと時間をかけて楽しく取り組んでいたことから、計画的な性格で細かい作業にも取り組むことができる特性を持っていると言えます。
    - **リーダーシップ**: プログラミングサークルの新歓活動を主導し、広報やイベントの計画・実行を行うなど、リーダーシップの要素が見受けられます。\n- **積極性**: 自分にとって不慣れなインスタグラムを使って活動の宣伝を行い、他のメンバーにも協力を促すなど、積極的に取り組んでいる姿勢が伺えます。\n- **コミュニケーション能力**: サークルのメンバーに協力を促したり、話し合いを通じてアイデアを出し合っていることから、コミュニケーション能力が高いと言えます。
    - **リーダーシップ**: 開発チームのリーダーを務め、進捗管理やメンバーのフォローを行い、チームをまとめる力を持っている。\n- **努力家**: 開発に真剣に取り組み、コンテストで金賞と最優秀賞を受賞することができた。努力を惜しまず、目標達成に向けて頑張る姿勢がある。\n- **技術力**: 高品質な制作物と効率的なプロジェクト運営を両立するために、技術力を身につける努力を行っている。
    」

    これらの情報から、特性を抽出または新しく作成し、50～100点で点数化します。
    特に点数が高いものから、必ず6つ表示してください。
    フォーマット例:「
    - **責任感**: __50__\n- **リーダーシップ**: __70__\n- **努力家**: __95__\n- **実験好き**: __84__\n- **チームワーク**: __72__\n- **コミュニケーション能力**: __92__
    」
  	`
	return openaiutil.FetchOpenAIResponse(prompt)
}
