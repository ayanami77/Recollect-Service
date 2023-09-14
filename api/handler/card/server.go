package card

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	"github.com/Seiya-Tagami/Recollect-Service/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type Handler interface {
	GetCard(c *gin.Context)
	ListCards(c *gin.Context)
	CreateCard(c *gin.Context)
	UpdateCard(c *gin.Context)
	DeleteCard(c *gin.Context)
}

type handler struct {
	cardInteractor card.Interactor
}

func New(cardInteractor card.Interactor) Handler {
	return &handler{cardInteractor}
}

func (h *handler) GetCard(c *gin.Context) {
	id := c.Param("id")

	card, err := h.cardInteractor.GetCard(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}

func (h *handler) ListCards(c *gin.Context) {
	tokenString, err := c.Cookie("user_token")
	if err != nil {
		panic(err)
	}
	token, err := utils.ParseToken(tokenString)
	var userID string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = claims["user_id"].(string)
	}

	cards, err := h.cardInteractor.ListCards(userID)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func (h *handler) CreateCard(c *gin.Context) {
	tokenString, err := c.Cookie("user_token")
	if err != nil {
		panic(err)
	}
	token, err := utils.ParseToken(tokenString)
	var userID string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = claims["user_id"].(string)
	}

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		panic(err)
	}
	cardReq.UserID = userID

	card, err := h.cardInteractor.CreateCard(cardReq)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}

func (h *handler) UpdateCard(c *gin.Context) {
	id := c.Param("id")
	tokenString, err := c.Cookie("user_token")
	if err != nil {
		panic(err)
	}
	token, err := utils.ParseToken(tokenString)
	var userID string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID = claims["user_id"].(string)
	}

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

	c.JSON(http.StatusOK, gin.H{"data": card})
}

func (h *handler) DeleteCard(c *gin.Context) {
	id := c.Param("id")

	err := h.cardInteractor.DeleteCard(id)
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusNoContent)
}
