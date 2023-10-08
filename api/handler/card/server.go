package card

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/response"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type Handler interface {
	ListCards(c *gin.Context)
	CreateCard(c *gin.Context)
	CreateCards(c *gin.Context)
	UpdateCard(c *gin.Context)
	DeleteCard(c *gin.Context)
}

type handler struct {
	cardInteractor card.Interactor
}

func New(cardInteractor card.Interactor) Handler {
	return &handler{cardInteractor}
}

func (h *handler) ListCards(c *gin.Context) {
	userID := userIDFromToken(c)

	cards, err := h.cardInteractor.ListCards(userID)
	if err != nil {
		panic(err)
	}

	cardsResponse := response.ToCardsResponse(&cards)

	c.JSON(http.StatusOK, gin.H{"data": cardsResponse})
}

func (h *handler) CreateCard(c *gin.Context) {
	userID := userIDFromToken(c)

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		panic(err)
	}
	cardReq.UserID = userID

	card, err := h.cardInteractor.CreateCard(cardReq)
	if err != nil {
		panic(err)
	}

	cardResponse := response.ToCardResponse(&card)

	c.JSON(http.StatusOK, gin.H{"data": cardResponse})
}

func (h *handler) CreateCards(c *gin.Context) {
	type BatchReq struct {
		Cards []entity.Card `json:"cards"`
	}
	var batchReq BatchReq

	userID := userIDFromToken(c)

	if err := c.BindJSON(&batchReq); err != nil {
		panic(err)
	}

	cardsReq := batchReq.Cards
	for i := range cardsReq {
		cardsReq[i].UserID = userID
	}

	cards, err := h.cardInteractor.CreateCards(cardsReq)
	if err != nil {
		panic(err)
	}

	cardsResponse := response.ToCardsResponse(&cards)

	c.JSON(http.StatusOK, gin.H{"data": cardsResponse})
}

func (h *handler) UpdateCard(c *gin.Context) {
	id := c.Param("id")
	userID := userIDFromToken(c)

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		panic(err)
	}

	cardReq.CardID = id
	cardReq.UserID = userID

	card, err := h.cardInteractor.UpdateCard(cardReq, id)
	if err != nil {
		panic(err)
	}

	cardResponse := response.ToCardResponse(&card)

	c.JSON(http.StatusOK, gin.H{"data": cardResponse})
}

func (h *handler) DeleteCard(c *gin.Context) {
	id := c.Param("id")

	err := h.cardInteractor.DeleteCard(id)
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusNoContent)
}

func userIDFromToken(c *gin.Context) string {
	tokenString, err := c.Cookie("user_token")
	if err != nil {
		panic(err)
	}
	token, err := utils.ParseToken(tokenString)
	var userID string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = claims["user_id"].(string)
	}
	return userID
}
