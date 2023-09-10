package card

import (
	"fmt"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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
	claims := jwt.ExtractClaims(c)
	userID := claims["user_id"]
	fmt.Printf("User %v", userID)

	cards, err := h.cardInteractor.ListCards(userID.(string))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func (h *handler) CreateCard(c *gin.Context) {
	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		panic(err)
	}

	card, err := h.cardInteractor.CreateCard(cardReq)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}

func (h *handler) UpdateCard(c *gin.Context) {
	id := c.Param("id")

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		panic(err)
	}

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
