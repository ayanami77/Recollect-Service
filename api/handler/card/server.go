package card

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/jwtutil"
	"github.com/Seiya-Tagami/Recollect-Service/api/handler/util/myerror"
	"github.com/Seiya-Tagami/Recollect-Service/api/response"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate mockgen -source=$GOFILE -destination=$GOPATH/Recollect-Service/api/mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE
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
