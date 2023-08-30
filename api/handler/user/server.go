package user

import (
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type handler struct {
	userInteractor user.Interactor
}

func New(userInteractor user.Interactor) Handler {
	return &handler{userInteractor}
}

func (h *handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	result, err := h.userInteractor.GetUser(id)
	if err != nil {
		panic(err)
	}

	print(result)
}

func (h *handler) CreateUser(c *gin.Context) {
	user := entity.User{}
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}

	result, err := h.userInteractor.CreateUser(user)
	if err != nil {
		panic(err)
	}

	print(result)
}

func (h *handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	user := entity.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(err)
	}

	result, err := h.userInteractor.UpdateUser(user, id)
	if err != nil {
		panic(err)
	}

	print(result)
}

func (h *handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.userInteractor.DeleteUser(id)
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusNoContent)
}
