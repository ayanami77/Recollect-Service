package card

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/jwtutil"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
	"github.com/Seiya-Tagami/Recollect-Service/api/response"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

//go:generate mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
type Handler interface {
	ListCards(c *gin.Context)
	CreateCard(c *gin.Context)
	CreateCards(c *gin.Context)
	UpdateCard(c *gin.Context)
	DeleteCard(c *gin.Context)
	UpdateAnalysisResult(c *gin.Context)
}

type handler struct {
	cardInteractor card.Interactor
}

func New(cardInteractor card.Interactor) Handler {
	return &handler{cardInteractor}
}

func (h *handler) ListCards(c *gin.Context) {
	sub, err := jwtutil.SubFromToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cards, err := h.cardInteractor.ListCards(sub)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardsResponse := response.ToCardsResponse(&cards)

	c.JSON(http.StatusOK, cardsResponse)
}

func (h *handler) CreateCard(c *gin.Context) {
	sub, err := jwtutil.SubFromToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		myerror.HandleError(c, err)
		return
	}
	cardReq.Sub = sub

	card, err := h.cardInteractor.CreateCard(cardReq)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardResponse := response.ToCardResponse(&card)

	c.JSON(http.StatusOK, cardResponse)
}

func (h *handler) CreateCards(c *gin.Context) {
	type BatchReq struct {
		Cards []entity.Card `json:"cards"`
	}
	var batchReq BatchReq

	sub, err := jwtutil.SubFromToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	if err := c.BindJSON(&batchReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardsReq := batchReq.Cards
	for i := range cardsReq {
		cardsReq[i].Sub = sub
	}

	cards, err := h.cardInteractor.CreateCards(cardsReq)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardsResponse := response.ToCardsResponse(&cards)

	c.JSON(http.StatusOK, cardsResponse)
}

func (h *handler) UpdateCard(c *gin.Context) {
	id := c.Param("id")
	sub, err := jwtutil.SubFromToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardReq.CardID = id
	cardReq.Sub = sub

	card, err := h.cardInteractor.UpdateCard(cardReq, id, sub)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardResponse := response.ToCardResponse(&card)

	c.JSON(http.StatusOK, cardResponse)
}

func (h *handler) DeleteCard(c *gin.Context) {
	id := c.Param("id")
	sub, err := jwtutil.SubFromToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	err = h.cardInteractor.DeleteCard(id, sub)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *handler) UpdateAnalysisResult(c *gin.Context) {
	id := c.Param("id")
	sub, err := subFromBearerToken(c)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardReq.CardID = id
	cardReq.Sub = sub

	card, err := h.cardInteractor.UpdateAnalysisResult(cardReq, id, sub)
	if err != nil {
		myerror.HandleError(c, err)
		return
	}

	cardResponse := response.ToCardResponse(&card)

	c.JSON(http.StatusOK, cardResponse)
}

func subFromBearerToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", myerror.InvalidRequest
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return "", myerror.InvalidRequest // Bearerトークンが見つからない場合
	}

	token, err := util.ParseToken(tokenString)
	if err != nil {
		return "", myerror.InvalidRequest
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, ok := claims["sub"].(string)
		if !ok {
			return "", myerror.InvalidRequest // subフィールドが存在しない場合
		}
		return sub, nil
	}

	return "", myerror.InvalidRequest // その他のエラー
}
