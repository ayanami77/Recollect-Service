package card

import (
	"fmt"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/card"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
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

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (h *handler) ListCards(c *gin.Context) {
	tokenString, err := c.Cookie("reco_cookie")
	if err != nil {
		panic(err)
	}
	token, err := ParseToken(tokenString)
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
	tokenString, err := c.Cookie("reco_cookie")
	if err != nil {
		panic(err)
	}
	token, err := ParseToken(tokenString)
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
