package card

import (
	"net/http"
	"strconv"

	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetCard(c *gin.Context)
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
	stringId := c.Param("id")

	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		panic(err)
	}

	card, err := h.cardInteractor.GetCard(uint(id))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
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
	stringId := c.Param("id")

	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		panic(err)
	}

	cardReq := entity.Card{}
	if err := c.BindJSON(&cardReq); err != nil {
		panic(err)
	}

	card, err := h.cardInteractor.UpdateCard(cardReq, uint(id))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}

func (h *handler) DeleteCard(c *gin.Context) {
	stringId := c.Param("id")

	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		panic(err)
	}

	err = h.cardInteractor.DeleteCard(uint(id))
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusNoContent)
}
